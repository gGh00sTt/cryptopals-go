package challenges

import (
	"fmt"
	"strings"
)

func Challenge16() {
	// TODO
	blocksize = 16
	keybyte, _ = generateRandomKey(blocksize)

	//create a string that will be encrypted
	input := "AAAAAAAAAAAAAAAAAAAAA;admin=true;aaaa"

	//encrypt the string
	ciphertext := encrypt_in_cbc(input)

	fmt.Println("Ciphertext: ", ciphertext)
	//find the blocksize and check if it matches the expected blocksize
	calculatedblocksize := find_blocksize()

	fmt.Println("Calculated blocksize: ", calculatedblocksize)

	prefix_len := find_prefix_length()

	fmt.Println("Prefix length: ", prefix_len)

	forcedCiphertext := cbcBitFlip()

	fmt.Println("Forced plaintext: ", Cbcdecrypt(string(forcedCiphertext), string(keybyte)))

	//decrypt the ciphertext and check if the admin=true is present
	admin := decrypt_and_check_admin(string(forcedCiphertext))

	if admin {
		println("Admin detected")
	} else {
		println("Admin not detected")
	}

}

func processAndRemoveChars(input string) string {
	input = strings.ReplaceAll(input, ";", "")
	input = strings.ReplaceAll(input, "=", "")

	result := "comment1=cooking%20MCs;userdata=" + input + ";comment2=%20like%20a%20pound%20of%20bacon"

	return result
}

func encrypt_in_cbc(input string) string {

	processedinput := processAndRemoveChars(input)

	plaintextbytes := []byte(processedinput)

	//determine the length of necessary padding
	paddinglength := blocksize - (len(plaintextbytes) % blocksize)

	//create padded plaintext
	plaintextbytes = padBytes(plaintextbytes, paddinglength)

	//perform cbcmode encryption

	ciphertext := Cbcencrypt(string(plaintextbytes), string(keybyte))

	return ciphertext

}

func find_blocksize() int {

	mystring := ""
	ciphertext := encrypt_in_cbc(mystring)
	ciphertextlength := len(ciphertext)
	newlen := ciphertextlength

	for newlen == ciphertextlength {
		mystring = mystring + "A"
		ciphertext = encrypt_in_cbc(mystring)
		newlen = len(ciphertext)
	}

	return newlen - ciphertextlength
}

func find_prefix_length() int {

	ciphertext_a := encrypt_in_cbc("A")
	ciphertext_b := encrypt_in_cbc("B")

	commonlength := 0

	for ciphertext_a[commonlength] == ciphertext_b[commonlength] {
		commonlength += 1
	}

	commonlength = commonlength / blocksize * blocksize

	for i := 1; i <= blocksize; i++ {
		ciphertext_a = encrypt_in_cbc(strings.Repeat("A", i) + "D")
		ciphertext_b = encrypt_in_cbc(strings.Repeat("A", i) + "E")

		if ciphertext_a[commonlength:commonlength+blocksize] == ciphertext_b[commonlength:commonlength+blocksize] {
			return commonlength + blocksize - i
		}
	}

	return 0

}

func decrypt_and_check_admin(ciphertext string) bool {

	plaintext := Cbcdecrypt(ciphertext, string(keybyte))

	return strings.Contains(plaintext, ";admin=true;")

}

func cbcBitFlip() []byte {
	// Get the length of a block and the length of the prefix
	blockLength := find_blocksize()
	prefixLength := find_prefix_length()

	// Compute the number of bytes to add to the prefix to make its length a multiple of blockLength
	additionalPrefixBytes := (blockLength - (prefixLength % blockLength)) % blockLength
	totalPrefixLength := prefixLength + additionalPrefixBytes

	// Compute the number of bytes to add to the plaintext to make its length a multiple of blockLength
	plaintext := "?admin?true"
	additionalPlaintextBytes := (blockLength - (len(plaintext) % blockLength)) % blockLength

	// Make the plaintext long one blockLength and encrypt it
	finalPlaintext := strings.Repeat("?", additionalPlaintextBytes) + plaintext
	ciphertext := encrypt_in_cbc(strings.Repeat("?", additionalPrefixBytes) + finalPlaintext)

	// Because XORing a byte with itself produces zero, we can produce the byte that we want
	// by changing the bytes of the block before the plaintext
	semicolon := ciphertext[totalPrefixLength-11] ^ '?' ^ ';'
	equals := ciphertext[totalPrefixLength-5] ^ '?' ^ '='

	// Put the pieces of our forged ciphertext together to generate the full ciphertext
	forcedCiphertext := ciphertext[:totalPrefixLength-11] + string(semicolon) + ciphertext[totalPrefixLength-10:totalPrefixLength-5] + string(equals) + ciphertext[totalPrefixLength-4:]
	return []byte(forcedCiphertext)
}

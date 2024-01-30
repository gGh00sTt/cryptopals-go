package challenges

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
)

var randomBytes []byte

func Challenge14() {
	randno, _ := generateRandomNumber(0, 256)

	randomBytes, _ = generateRandomKey(randno)

	keybyte, _ = generateRandomKey(16)

	byte_at_a_time_decryption()

}

func harder_encryption_oracle(plaintext string, key string) string {
	plaintextbytes := []byte(plaintext)

	keybytes := []byte(key)

	unknown := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"

	unknowndec, _ := b64.StdEncoding.DecodeString(unknown)

	plaintextbytespadded := []byte{}

	plaintextbytespadded = append(plaintextbytespadded, randomBytes...)
	plaintextbytespadded = append(plaintextbytespadded, plaintextbytes...)
	plaintextbytespadded = append(plaintextbytespadded, unknowndec...)

	ciphertext := Ecbencrypt(string(plaintextbytespadded), string(keybytes))

	return ciphertext
}

func find_blocksize_used() int {
	mystring := ""
	ciphertext := harder_encryption_oracle(mystring, string(keybyte))
	ciphertextlength := len(ciphertext)
	newlen := ciphertextlength

	for newlen == ciphertextlength {
		mystring = mystring + "A"
		ciphertext = harder_encryption_oracle(mystring, string(keybyte))
		newlen = len(ciphertext)
	}

	return newlen - ciphertextlength
}

func has_equal_block(ciphertext string, blocksize int) bool {
	for i := 0; i < len(ciphertext)-2*blocksize; i += blocksize {
		if bytes.Equal([]byte(ciphertext[i:i+blocksize]), []byte(ciphertext[i+blocksize:i+2*blocksize])) {
			return true
		}
	}

	return false
}

func find_prefix_len(blockLength int) int {
	ciphertext1 := []byte(harder_encryption_oracle("", string(keybyte)))
	ciphertext2 := []byte(harder_encryption_oracle("a", string(keybyte)))

	var prefixLength int
	for i := 0; i < len(ciphertext2); i += blockLength {
		if !bytes.Equal(ciphertext1[i:i+blockLength], ciphertext2[i:i+blockLength]) {
			prefixLength = i
			break
		}
	}

	for i := 0; i < blockLength; i++ {
		fakeInput := bytes.Repeat([]byte{0}, 2*blockLength+i)
		ciphertext := harder_encryption_oracle(string(fakeInput), string(keybyte))

		if has_equal_block(ciphertext, blockLength) {
			return prefixLength + blockLength - i
		}
	}

	return prefixLength
}

func get_next_byte(prefixLength int, blockLength int, currDecryptedMessage []byte) byte {
	lengthToUse := (blockLength - prefixLength - (1 + len(currDecryptedMessage))) % blockLength
	if lengthToUse < 0 {
		lengthToUse += blockLength
	}

	myInput := bytes.Repeat([]byte{'A'}, lengthToUse)

	crackingLength := prefixLength + lengthToUse + len(currDecryptedMessage) + 1

	realCiphertext := []byte(harder_encryption_oracle(string(myInput), string(keybyte)))

	for i := 0; i < 256; i++ {
		fakeCiphertext := []byte(harder_encryption_oracle(string(append(myInput, append(currDecryptedMessage, byte(i))...)), string(keybyte)))

		if bytes.Equal(fakeCiphertext[:crackingLength], realCiphertext[:crackingLength]) {
			return byte(i)
		}
	}

	return 0
}

func byte_at_a_time_decryption() {

	block_length := find_blocksize_used()

	plaintexttest := string(bytes.Repeat([]byte("s"), block_length*6))

	cipertexttest := harder_encryption_oracle(plaintexttest, string(keybyte))

	mode := DetectEncryptionMode(cipertexttest)

	println("encryption mode :", mode)

	prefix_len := find_prefix_len(block_length)
	secret_len := len(harder_encryption_oracle("", string(keybyte))) - prefix_len

	fmt.Println("prefix length :", prefix_len)
	fmt.Println("secret length :", secret_len)

	secret_padding := []byte{}

	for i := 0; i < secret_len; i++ {
		secret_padding = append(secret_padding, get_next_byte(prefix_len, block_length, secret_padding))
	}

	fmt.Println(string(secret_padding))

}

package challenges

import (
	"crypto/aes"
	b64 "encoding/base64"
	"fmt"
	"os"
)

func Challenge10() {
	key := "YELLOW SUBMARINE"
	// plaintext := "Hello I am csangam"

	// fmt.Printf("plaintext : %s \n", plaintext)

	// ciphertext := cbcencrypt(plaintext, key)

	file, err := os.Open("challenges/challenge10.txt")

	if err != nil {
		panic(err)
	}

	fileinfo, _ := file.Stat()

	filesize := fileinfo.Size()

	filebytes := make([]byte, filesize)

	file.Read(filebytes)
	filebytesdec, _ := b64.StdEncoding.DecodeString(string(filebytes))

	ciphertext := string(filebytesdec)

	fmt.Printf("ciphertext : %s \n", ciphertext)

	plaintextafter := cbcdecrypt(ciphertext, key)

	fmt.Printf("plaintext : %s \n", plaintextafter)

}

func encryptAesBlock(key []byte, plaintextblock []byte) []byte {
	c, _ := aes.NewCipher(key)

	ciphertextblock := make([]byte, len(plaintextblock))

	c.Encrypt(ciphertextblock, plaintextblock)

	return ciphertextblock
}

func decryptAesBlock(key []byte, ciphertextblock []byte) []byte {
	c, _ := aes.NewCipher(key)

	plaintextblock := make([]byte, len(ciphertextblock))

	c.Decrypt(plaintextblock, ciphertextblock)

	return plaintextblock
}

func createIV(keysize int) []byte {
	ivbyte := []byte{}

	for i := 0; i < keysize; i++ {
		ivbyte = append(ivbyte, byte(0))
	}

	return ivbyte
}

func xor(block1 []byte, block2 []byte) []byte {
	xorblock := make([]byte, len(block1))
	for i := 0; i < len(block1); i++ {
		xorblock[i] = block1[i] ^ block2[i]
	}

	return xorblock
}

func createblocks(plaintextblock []byte, keysize int) [][]byte {
	byteBlockArray := [][]byte{}
	for i := 0; i < len(plaintextblock); i = i + keysize {

		byteBlockArray = append(byteBlockArray, plaintextblock[i:i+keysize])
	}

	return byteBlockArray
}

func addpadding(plaintextbyte []byte, keysize int) []byte {

	for len(plaintextbyte)%keysize != 0 {
		plaintextbyte = append(plaintextbyte, byte(0))
	}

	return plaintextbyte
}

func cbcencrypt(plaintext string, key string) string {
	keybyte := []byte(key)
	plaintextbyte := []byte(plaintext)

	keysize := len(keybyte)

	plaintextbyte = addpadding(plaintextbyte, keysize)

	ivblock := createIV(keysize)

	plaintextbyteblocks := createblocks(plaintextbyte, keysize)

	xorblock := ivblock

	ciphertextbyte := []byte{}
	for i := 0; i < len(plaintextbyteblocks); i++ {

		xoredblock := xor(plaintextbyteblocks[i], xorblock)

		cipherblock := encryptAesBlock(keybyte, xoredblock)

		ciphertextbyte = append(ciphertextbyte, cipherblock...)

		xorblock = cipherblock
	}

	return string(ciphertextbyte)

}

func cbcdecrypt(ciphertext string, key string) string {
	keybyte := []byte(key)
	ciphertextbyte := []byte(ciphertext)

	keysize := len(keybyte)

	ivblock := createIV(keysize)

	ciphertextbyteblocks := createblocks(ciphertextbyte, keysize)

	xorblock := ivblock

	plaintextbyte := []byte{}
	for i := 0; i < len(ciphertextbyteblocks); i++ {

		plaintextblock := decryptAesBlock(keybyte, ciphertextbyteblocks[i])

		xoredblock := xor(plaintextblock, xorblock)

		plaintextbyte = append(plaintextbyte, xoredblock...)

		xorblock = ciphertextbyteblocks[i]
	}

	return string(plaintextbyte)

}

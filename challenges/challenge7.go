package challenges

import (
	"crypto/aes"
	b64 "encoding/base64"
	"os"
)

func Challenge7() {

	file, err := os.Open("challenges/challenge7.txt")

	if err != nil {
		panic(err)
	}

	fileInfo, _ := file.Stat()

	fileSize := fileInfo.Size()

	fcontent := make([]byte, fileSize)

	file.Read(fcontent)

	fcontentdec, _ := b64.StdEncoding.DecodeString(string(fcontent))

	key := []byte("YELLOW SUBMARINE")

	fcontentdecdecrypted := decryptAes(key, fcontentdec)

	println(string(fcontentdecdecrypted))

}

func decryptAes(key []byte, cipher []byte) []byte {
	c, _ := aes.NewCipher(key)

	keysize := len(key)

	byteBlockArray := [][]byte{}
	for i := 0; i < len(cipher); i = i + keysize {

		byteBlockArray = append(byteBlockArray, cipher[i:i+keysize])
	}

	decrypted := [][]byte{}

	for _, block := range byteBlockArray {

		decblock := make([]byte, len(block))
		c.Decrypt(decblock, block)
		decrypted = append(decrypted, decblock)
	}

	concdecrypted := make([]byte, len(decrypted)*keysize)

	for i := 0; i < len(decrypted); i++ {
		for j := 0; j < len(decrypted[i]); j++ {
			concdecrypted[i*len(decrypted[j])+j] = decrypted[i][j]
		}
	}

	return concdecrypted

}

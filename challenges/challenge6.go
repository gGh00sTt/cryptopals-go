package challenges

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Challenge6() {

	file, err := os.Open("challenges/challenge6.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		panic(err)
	}

	lenFile := fileInfo.Size()

	fbytes := make([]byte, lenFile)

	_, err = file.Read(fbytes)

	if err != nil {
		panic(err)
	}

	fb64DecBytes, errn := b64.StdEncoding.DecodeString(string(fbytes))

	if errn != nil {
		panic(errn)
	}

	keysize := getKeySize(fb64DecBytes)

	//padding extra bits
	for len(fb64DecBytes)%keysize != 0 {
		fb64DecBytes = append(fb64DecBytes, 0)
	}

	byteBlockArray := [][]byte{}
	for i := 0; i < len(fb64DecBytes); i = i + keysize {

		byteBlockArray = append(byteBlockArray, fb64DecBytes[i:i+keysize])
	}

	checkBlockArray := [][]byte{}

	for i := 0; i < keysize; i++ {
		singleCheckBlockByte := []byte{}
		for j := 0; j < len(byteBlockArray); j++ {
			singleCheckBlockByte = append(singleCheckBlockByte, byteBlockArray[j][i])
		}
		checkBlockArray = append(checkBlockArray, singleCheckBlockByte)
	}

	key := ""

	for i := 0; i < len(checkBlockArray); i++ {

		convbyte := checkBlockArray[i]

		key += string(singleCharXor(convbyte))

	}

	decrypted := ""

	for _, block := range byteBlockArray {

		for i, ibyte := range block {
			decrypted += string(ibyte ^ byte(key[i]))
		}
	}

	fmt.Println("key: ", key)
	fmt.Println(decrypted)

}

func getKeySize(decodedBytes []byte) int {
	score := 100000.0
	smallestKeySize := 2
	noOfSamples := 10
	for keysize := 2; keysize < 41; keysize++ {
		sum := 0
		for i := 0; i < noOfSamples; i++ {
			dist, _ := hamming(string(decodedBytes[i*keysize:i*keysize+keysize]), string(decodedBytes[i*keysize+keysize:i*keysize+keysize*2]))
			sum += dist
		}

		editDist := float64(sum) / float64(keysize)
		avgEditDist := editDist / float64(noOfSamples)

		if avgEditDist < score {
			smallestKeySize = keysize
			score = avgEditDist
		}
	}

	return smallestKeySize

}

func hamming(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("strings are unequal in length")
	}

	byte1 := []byte(a)
	byte2 := []byte(b)
	diff := 0

	for i := 0; i < len(byte1); i++ {
		b1 := byte1[i]
		b2 := byte2[i]

		for j := 0; j < 8; j++ {
			b1bit := (b1 >> j) & 1
			c1bit := (b2 >> j) & 1

			if b1bit != c1bit {
				diff += 1
			}
		}
	}

	return diff, nil
}

func singleCharXor(convbyte []byte) byte {
	score := 0
	keyChar := byte(0)
	freqorder := " etaoinshrdlcumwfgypbvkjxqz"
	freqorderrev := ""
	for i := len(freqorder) - 1; i >= 0; i-- {
		freqorderrev += string(freqorder[i])
	}

	for char := byte(0); char < 255; char++ {
		newbyte := make([]byte, len(convbyte))
		for i := 0; i < len(convbyte); i++ {
			newbyte[i] = convbyte[i] ^ byte(char)
		}

		newscore := 0

		for _, char := range newbyte {
			if char > 64 && char < 127 || char > 96 && char < 123 || char == 32 {

				position := strings.Index(freqorderrev, strings.ToLower(string(char)))
				newscore += position
			}
		}

		if newscore >= score {
			score = newscore
			keyChar = char
		}

	}

	return keyChar

}

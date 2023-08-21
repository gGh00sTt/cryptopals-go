package challenges

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func Challenge3() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	convbyte, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	englishLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	freqorder := "etaoinshrdlcumwfgypbvkjxqz"
	freqorderrev := ""
	for i := len(freqorder) - 1; i >= 0; i-- {
		freqorderrev += string(freqorder[i])
	}

	score := 0
	finalbyte := make([]byte, len(convbyte))
	xorcharacter := ""

	for _, char := range englishLetters {
		newbyte := make([]byte, len(convbyte))
		for i := 0; i < len(convbyte); i++ {
			newbyte[i] = convbyte[i] ^ byte(char)
		}

		newscore := 0

		for _, char := range newbyte {
			if char > 64 && char < 127 || char > 96 && char < 123 {

				position := strings.Index(freqorderrev, strings.ToLower(string(char)))
				newscore += position
			}
		}

		if newscore >= score {
			score = newscore
			finalbyte = newbyte
			xorcharacter = string(char)
		}
	}

	fmt.Println("Decrypted String : ", string(finalbyte))
	fmt.Println("Xored Character : ", xorcharacter)

}

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

	freqorder := " etaoinshrdlcumwfgypbvkjxqz"
	freqorderrev := ""
	for i := len(freqorder) - 1; i >= 0; i-- {
		freqorderrev += string(freqorder[i])
	}

	score := 0
	finalbyte := make([]byte, len(convbyte))
	xorcharacter := byte(0)

	for char := byte(0); char < 255; char++ {
		newbyte := make([]byte, len(convbyte))
		for i := 0; i < len(convbyte); i++ {
			newbyte[i] = convbyte[i] ^ char
		}

		newscore := 0

		for _, char := range newbyte {
			if char > 0 && char < 255 {

				position := strings.Index(freqorderrev, string(char))
				newscore += position
			}
		}

		if newscore >= score {
			score = newscore
			finalbyte = newbyte
			xorcharacter = char
		}
	}

	fmt.Println("Decrypted String : ", string(finalbyte))
	fmt.Println("Xored Character : ", string(xorcharacter))
	fmt.Println("Score : ", score)
}

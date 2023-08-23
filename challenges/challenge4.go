package challenges

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func Challenge4() {
	file, err := os.Open("challenges/challenge4.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	finalWord := ""
	for scanner.Scan() {
		input := scanner.Text()

		convbyte, err := hex.DecodeString(input)
		if err != nil {
			panic(err)
		}

		freqorder := " etaoinshrdlcumwfgypbvkjxqz"
		freqorderrev := ""
		for i := len(freqorder) - 1; i >= 0; i-- {
			freqorderrev += string(freqorder[i])
		}

		//score := 0
		//finalbyte := make([]byte, len(convbyte))
		//xorcharacter := ""

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
				finalWord = string(newbyte)
			}

		}

	}
	fmt.Println("Decrypted String : ", finalWord)
	fmt.Println("Score : ", score)

}

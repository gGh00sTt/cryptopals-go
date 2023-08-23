package challenges

import (
	"bufio"
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

	count := 0

	for scanner.Scan() {
		count++
		input := scanner.Text()

		//string to bytes

		convbyte := []byte(input)
		if err != nil {
			panic(err)
		}

		englishLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		freqorder := "etaoinshrdlcumwfgypbvkjxqz"
		freqorderrev := ""
		for i := len(freqorder) - 1; i >= 0; i-- {
			freqorderrev += string(freqorder[i])
		}

		//score := 0
		//finalbyte := make([]byte, len(convbyte))
		//xorcharacter := ""

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
				//finalbyte = newbyte
				//xorcharacter = string(char)
				fmt.Println("Decrypted String : ", string(newbyte))
				fmt.Println("Score : ", score)
			}

		}
		//fmt.Println("Count : ", count)
		//fmt.Println("line : ", input)

		//fmt.Println("Xored Character : ", xorcharacter)
	}

}

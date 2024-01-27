package challenges

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func Challenge8() {

	filePath := "challenges/challenge8.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineBytes := []byte(scanner.Text())
		decLine, _ := hex.DecodeString(string(lineBytes))
		lines = append(lines, decLine)
	}

	lowest := 10000
	final := ""
	finalline := 0

	for i, line := range lines {

		sum := getDiff(line)

		if sum < lowest {
			lowest = sum
			final = hex.EncodeToString(line)
			finalline = i + 1

		}
	}

	println(finalline, final)

}

func getDiff(decodedBytes []byte) int {
	keysize := 16
	sum := 0
	noOfSamples := len(decodedBytes) / keysize

	for i := 0; i < noOfSamples-1; i++ {
		for j := i + 1; j < noOfSamples; j++ {
			dist, _ := hamming(string(decodedBytes[i*keysize:i*keysize+keysize]), string(decodedBytes[j*keysize:j*keysize+keysize]))
			sum += dist
		}

	}

	return sum

}

package challenges

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Challenge11() {

	plaintext := "Mama, take this badge off of me, I can't use it anymore, It's gettin' dark, too dark for me to see, I feel like I'm knockin' on heaven's door. I'm knockin' on heaven's door. Knock, knock, knockin' on heaven's door Knock, knock, knockin' on heaven's door Knock, knock, knockin' on heaven's door Knock, knock, knockin' on heaven's door"

	encryption_oracle(plaintext)
}

func generateRandomKey(length int) ([]byte, error) {
	if length <= 0 {
		length = 16
	}

	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func generateRandomBit() (int, error) {
	randomBit, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return 0, err
	}
	return int(randomBit.Int64()), nil
}

func generateRandomNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("invalid range: min is greater than max")
	}

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return int(randomNumber.Int64()) + min, nil
}

func encryption_oracle(plaintext string) {
	plaintextbytes := []byte(plaintext)
	randno, _ := generateRandomNumber(5, 10)

	keybytes, _ := generateRandomKey(16)

	plaintextbytespadded := []byte{}

	padding := []byte{}

	for i := 0; i < randno; i++ {
		padding = append(padding, byte(randno))
	}

	plaintextbytespadded = append(plaintextbytespadded, padding...)
	plaintextbytespadded = append(plaintextbytespadded, plaintextbytes...)
	plaintextbytespadded = append(plaintextbytespadded, padding...)

	choice, _ := generateRandomBit()

	ciphertext := ""

	print("Encryption mode : ")
	if choice == 0 {
		println("ecb")
		ciphertext = Ecbencrypt(string(plaintextbytespadded), string(keybytes))
	} else {
		println("cbc")
		ciphertext = Cbcencrypt(string(plaintextbytespadded), string(keybytes))
	}

	println("Detected mode : ", DetectEncryptionMode(ciphertext))

}

func DetectEncryptionMode(ciphertext string) string {
	mode := ""
	diff := countAESecbRepetitions([]byte(ciphertext))

	if diff == 0 {
		mode = "cbc"
	} else {
		mode = "ecb"
	}

	return mode

}

func countAESecbRepetitions(ciphertext []byte) int {
	blocksize := 16
	chunks := make(map[string]int)

	for i := 0; i < len(ciphertext); i += blocksize {
		chunk := ciphertext[i : i+blocksize]
		chunks[string(chunk)]++
	}

	numberofDuplicates := 0
	for _, count := range chunks {
		if count > 1 {
			numberofDuplicates += count - 1
		}
	}

	return numberofDuplicates
}

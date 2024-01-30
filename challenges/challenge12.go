package challenges

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
)

func Challenge12() {

	keybytes := []byte("sangamchaulagain")
	key := string(keybytes)
	plaintext := "My name is sangam"

	//println(len(ciphertext))
	//println(len(b64.StdEncoding.EncodeToString([]byte(ciphertext))))

	//detect block size
	plaintextsingle := string([]byte(plaintext)[0])

	current_size := len(encryption_oracle_ecb(plaintextsingle, key))
	block_size := 0

	for i := 0; i < 256; i++ {
		plaintextrepeat := bytes.Repeat([]byte(plaintextsingle), i)

		new_size := len(encryption_oracle_ecb(string(plaintextrepeat), key))

		if new_size != current_size {
			block_size = new_size - current_size
			break
		}

	}

	//detect cipher mode (ecb)
	plaintexttest := string(bytes.Repeat([]byte("s"), block_size*2))

	cipertexttest := encryption_oracle_ecb(plaintexttest, key)

	mode := DetectEncryptionMode(cipertexttest)

	println("encryption mode :", mode)

	//performing dictionary attack to find unknown text

	attackDict := make(map[string]string)

	count := 1
	start := 0
	end := 16
	lastchar := ""
	maxlen := len(encryption_oracle_ecb("", key))

	for count < maxlen {

		for i := 0; i < 128; i++ {
			attack := ""
			character := fmt.Sprintf("%c", i)
			reqattackTextLen := count * 16
			attack = string(bytes.Repeat([]byte("s"), reqattackTextLen-len(lastchar)-1)) + lastchar + character
			attackText := attack

			attackcipertext := encryption_oracle_ecb(attackText, key)

			attackDict[attackcipertext[start:end]] = character
		}

		reqattackTextLen := count * 16
		attackTxt := string(bytes.Repeat([]byte("s"), reqattackTextLen-len(lastchar)-1))
		attackcipherText := encryption_oracle_ecb(attackTxt, key)

		lastchar += attackDict[attackcipherText[start:end]]

		end += 16
		start += 16
		count += 1
		attackDict = make(map[string]string)

	}

	fmt.Println(lastchar)

}

func encryption_oracle_ecb(plaintext string, key string) string {
	plaintextbytes := []byte(plaintext)

	keybytes := []byte(key)

	unknown := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"

	unknowndec, _ := b64.StdEncoding.DecodeString(unknown)

	plaintextbytespadded := []byte{}

	plaintextbytespadded = append(plaintextbytespadded, plaintextbytes...)
	plaintextbytespadded = append(plaintextbytespadded, unknowndec...)

	ciphertext := Ecbencrypt(string(plaintextbytespadded), string(keybytes))

	return ciphertext
}

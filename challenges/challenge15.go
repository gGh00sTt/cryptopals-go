package challenges

import (
	"fmt"
	"strings"
)

func Challenge15() {
	text1 := "ICE ICE BABY\x04\x04\x04\x04"
	text2 := "ICE ICE BABY\x05\x05\x05\x05"
	text3 := "ICE ICE BABY\x01\x02\x03\x04"
	text4 := "ICE ICE BABY\x05\x05\x05\x05\x05"

	fmt.Println(isValidPadding(text1))
	fmt.Println(isValidPadding(text2))
	fmt.Println(isValidPadding(text3))
	fmt.Println(isValidPadding(text4))
}

func isValidPadding(plaintextpadded string) bool {
	required := len(plaintextpadded)
	plaintextpadded = fmt.Sprintf("%q", plaintextpadded)

	plaintextpadded = plaintextpadded[1 : len(plaintextpadded)-1]
	plaintextpaddedsplits := strings.Split(plaintextpadded, "\\x")

	plaintext := plaintextpaddedsplits[0]

	plaintextpaddedagain := string(padBytes([]byte(plaintext), required))

	plaintextpaddedagainformatted := fmt.Sprintf("%q", plaintextpaddedagain)
	plaintextpaddedagainformatted = plaintextpaddedagainformatted[1 : len(plaintextpaddedagainformatted)-1]

	if plaintextpaddedagainformatted == plaintextpadded {
		return true
	} else {
		return false
	}

}

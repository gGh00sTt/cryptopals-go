package challenges

import (
	"fmt"
)

func Challenge5(){
	myString := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	newString := RepeatingKeyXOR(myString, key)
	fmt.Printf("%x", newString)
	fmt.Println()

}

func RepeatingKeyXOR(myString string, key string) string {
	var result string
	for i := 0; i < len(myString); i++ {
		result += string(myString[i] ^ key[i%len(key)])
	}
	return result
}
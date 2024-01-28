package challenges

import "fmt"

func Challenge9() {
	mystr := "YELLOW SUBMARINE"
	mystrbytes := []byte(mystr)
	mystrbytespadded := padBytes(mystrbytes, 20)
	println("before size : ", len(mystrbytes))
	fmt.Printf("before raw form : %v\n", mystrbytes)
	println("after size : ", len(mystrbytespadded))
	fmt.Printf("before raw form : %v\n", mystrbytespadded)
}

func padBytes(plaintext []byte, sizewithpadding int) []byte {
	required := sizewithpadding - len(plaintext)
	for i := 0; i < required; i++ {
		plaintext = append(plaintext, byte(required))
	}
	return plaintext
}

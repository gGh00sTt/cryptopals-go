package challenges

import (
	"encoding/base64"
	"encoding/hex"
)

func Challenge1() {

	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	convbyte, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	convbase64 := base64.StdEncoding.EncodeToString(convbyte)
	println(convbase64)

}

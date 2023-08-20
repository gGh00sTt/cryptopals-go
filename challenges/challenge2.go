package challenges

import (
	"encoding/hex"
	"fmt"
)

func Challenge2() {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"

	convbyte1, err := hex.DecodeString(hex1)
	if err != nil {
		panic(err)
	}
	convbyte2, err := hex.DecodeString(hex2)
	if err != nil {
		panic(err)
	}

	// XOR the two byte arrays
	for i := 0; i < len(convbyte1); i++ {
		convbyte1[i] = convbyte1[i] ^ convbyte2[i]
	}

	// Convert the XOR'd byte array to hex
	convhex := hex.EncodeToString(convbyte1)
	fmt.Println(convhex)

}

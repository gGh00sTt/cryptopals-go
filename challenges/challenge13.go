package challenges

import (
	"fmt"
	"math/rand"
	"strings"
)

var keybyte []byte
var blocksize int

func Challenge13() {
	blocksize = 16

	keybyte, _ = generateRandomKey(blocksize)

	email_header_len := len("email=")
	required_padding_first_block := blocksize - (email_header_len % blocksize)
	keyword := "admin"
	required_padding_second_block := blocksize - (len(keyword) % blocksize)
	padded_email_first_block := strings.Repeat(string(byte(required_padding_first_block)), required_padding_first_block)
	padded_email_second_block := strings.Repeat(string(byte(required_padding_second_block)), required_padding_second_block)
	email_input_1 := padded_email_first_block + keyword + padded_email_second_block

	//crafitng email_input_2 to get the admin role
	email_input_2 := "san@gmail.com"

	email1_enc := profile_for(email_input_1)
	email2_enc := profile_for(email_input_2)

	email2_dec := decrypt_profile(email2_enc)
	email2_dec_split := strings.Split(email2_dec, "user")
	position_to_replace := len(email2_dec_split[0])

	final_email_enc := email2_enc[:position_to_replace] + email1_enc[blocksize:blocksize*2]
	final_email_dec := decrypt_profile(final_email_enc)

	fmt.Println(final_email_dec)

	final_email_dict := parse(final_email_dec)

	for key, value := range final_email_dict {
		fmt.Println(key, value)
	}

}

func parse(accinfo string) map[string]string {
	accdict := make(map[string]string)

	accinfosplitted := strings.Split(accinfo, "&")

	for _, value := range accinfosplitted {
		iinfo := strings.Split(value, "=")

		accdict[iinfo[0]] = iinfo[1]
	}

	return accdict

}

func profile_for(email string) string {
	uid := rand.Intn(100)
	role := "user"

	if strings.Contains(email, "&") || strings.Contains(email, "=") {
		panic("email cannot contain & or =")
	}

	encoded_profile := fmt.Sprintf("email=%s&uid=%d&role=%s", email, uid, role)

	ciphertext := Ecbencrypt(encoded_profile, string(keybyte))

	return ciphertext
}

func decrypt_profile(cipertext string) string {
	plaintext := Ecbdecrypt(cipertext, string(keybyte))

	return plaintext
}

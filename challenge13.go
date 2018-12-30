package main

import "fmt"
import "strings"

import "crypto/aes"

//compiles with go run challenge13.go challenge11.go

func kEqualsVParsing(str string) map[string]string {
	m := make(map[string]string)
	temp := strings.Split(str, "&")
	for i := range temp {
		kv := strings.Split(temp[i], "=")
		m[kv[0]] = kv[1]
	}
	return m
}

func profile_for(email string) string {
	email = strings.Replace(email, "&", "", -1)
	email = strings.Replace(email, "=", "", -1)
	return "email=" + email + "&uid=10&role=user"
}

func AESDecrypt(cipherIn []byte, key []byte) []byte {
	cipher := make([]byte, len(cipherIn))
	copy(cipher, cipherIn)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// dst := make([]byte, 16)
	for i := 0; i < len(cipher); i += 16 {
		block.Decrypt(cipher[i:i+16], cipher[i:i+16])
	}
	return cipher
}

func encryptProfile(email string, key []byte) []byte {
	b := []byte(profile_for(email))
	if len(b)%16 != 0 {
		padding := make([]byte, 16-len(b)%16)
		b = append(b, padding...)
	}

	return ecbEncrypt(key, b)
}

func main() {

	email := "AAAAAAAAAAAAA"
	key := randomKey()

	encrypted1 := encryptProfile(email, key)

	fmt.Println(len(encrypted1))

	email = "AAAAAAAAAAadminAAAAAAAAAAA"
	email = strings.Replace(email, "A", string(0x00), -1)
	encrypted2 := encryptProfile(email, key)

	encryptedDef := append(encrypted1[:32], encrypted2[16:32]...)
	decrypted := AESDecrypt(encryptedDef, key)
	fmt.Println(string(decrypted))
}

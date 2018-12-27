package main

import "fmt"

import hex "encoding/hex"

func xor_equal_len(hex_s1 string, hex_s2 string) string {
	bytes_s1, _ := hex.DecodeString(hex_s1)
	bytes_s2, _ := hex.DecodeString(hex_s2)
	bytes_xor := make([]byte, len(bytes_s1))
	for i := 0; i < len(bytes_s1); i++ {
		bytes_xor[i] = bytes_s1[i] ^ bytes_s2[i]
	}
	return hex.EncodeToString(bytes_xor)
}

func main() {
	fmt.Println("challenge 2 : ", xor_equal_len("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
}

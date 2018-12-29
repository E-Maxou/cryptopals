package main

import "fmt"

import b64 "encoding/base64"
import hex "encoding/hex"

func hex_to_base64(hex_s string) string {
	bytes, _ := hex.DecodeString(hex_s)
	sEnc := b64.StdEncoding.EncodeToString(bytes)
	return sEnc
}

func main() {
	fmt.Println("challenge 1 : ", hex_to_base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}

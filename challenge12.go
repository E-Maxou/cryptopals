package main

// compiles with go run challene12.go challenge11.go

import "fmt"
import b64 "encoding/base64"
import hex "encoding/hex"

func main() {
	b64string := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"
	byteStr, _ := b64.StdEncoding.DecodeString(b64string)
	hexStr := hex.EncodeToString(byteStr)
	encrypted := encryptRandom("414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141" + hexStr)
	fmt.Println(oracleMode(encrypted))
}

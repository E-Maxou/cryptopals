package main

// compiles with go run challene12.go challenge11.go, you need to uncomment main

import "fmt"
import b64 "encoding/base64"
import hex "encoding/hex"

func truncatedHexUnk(nbToRm int) string {
	b64Unk := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"
	byteUnk, _ := b64.StdEncoding.DecodeString(b64Unk)
	hexUnk := hex.EncodeToString(byteUnk)
	if nbToRm*2 > len(hexUnk) {
		return ""
	}
	return hexUnk[nbToRm*2:]
}

func encryptConcatenation(myHexString string, hexUnk string, key []byte) []byte {
	concat, _ := hex.DecodeString(myHexString + hexUnk)
	return ecbEncrypt(key, concat)
}

func discoverKeysize() int {
	for i := 1; i < 65; i++ {
		myString := ""
		for j := 0; j < i; j++ {
			myString += "FF"
		}
		key := randomKey()
		encrypted := encryptConcatenation(myString, truncatedHexUnk(0), key)

		if oracleMode(encrypted) == "ecb" {
			// dividing by 2 because we fed exactly two blocks here (we return the first time we meet two identical blocks)
			return i / 2
		}
	}
	return -1
}

func makeByteMap(key []byte) map[string]int {
	m := make(map[string]int)
	beginning := ""
	for i := 0; i < 15; i++ {
		beginning += "41"
	}
	for i := 0; i < 256; i++ {
		blockHex := beginning
		lastByte := fmt.Sprintf("%x", i)
		if len(lastByte) == 0 {
			lastByte = "00"
		}
		if len(lastByte) == 1 {
			lastByte = "0" + lastByte
		}
		blockHex += lastByte
		block, _ := hex.DecodeString(blockHex)
		encrypted := ecbEncrypt(key, block)
		encryptedHex := hex.EncodeToString(encrypted)
		m[encryptedHex] = i
	}
	return m
}

func ecbOracle(key []byte) string {
	m := makeByteMap(key)
	cat := ""
	for i := 0; i < discoverKeysize()-1; i++ {
		cat += "41"
	}
	result := make([]byte, 0)
	for i := 0; truncatedHexUnk(i) != ""; i++ {
		toEncrypt, _ := hex.DecodeString(cat + truncatedHexUnk(i)[:2])
		encrypted := ecbEncrypt(key, toEncrypt)
		result = append(result, byte(m[hex.EncodeToString(encrypted)]))
	}
	return string(result)
}

func main() {
	key := randomKey()
	fmt.Print("Challenge 12 : ", ecbOracle(key))
}

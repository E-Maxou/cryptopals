package main

// compiles with go run challene12.go challenge11.go, you need to uncomment main

import "fmt"
import b64 "encoding/base64"
import hex "encoding/hex"

func hexUnk() string {
	b64Unk := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"
	byteUnk, _ := b64.StdEncoding.DecodeString(b64Unk)
	hexUnk := hex.EncodeToString(byteUnk)
	return hexUnk
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
		encrypted := encryptConcatenation(myString, hexUnk(), key)

		if oracleMode(encrypted) == "ecb" {
			// dividing by 2 because we fed exactly two blocks here (we return the first time we meet two identical blocks)
			return i / 2
		}
	}
	return -1
}

func makeByteMap(key []byte, beginning string) map[string]int {
	m := make(map[string]int)
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

// currently this only works for the first block
func decryptFirstBlock(key []byte) []byte {
	previousBlockHex := "414141414141414141414141414141"
	result := make([]byte, 0)
	block := make([]byte, 0)
	current := 0
	for j := 0; j < 16; j++ {
		known := hex.EncodeToString(result)
		m := makeByteMap(key, previousBlockHex+known)

		toEncrypt, _ := hex.DecodeString(previousBlockHex + hexUnk())

		encrypted := ecbEncrypt(key, toEncrypt)
		block = append(block, byte(m[hex.EncodeToString(encrypted[current:current+16])]))
		result = append(result, byte(m[hex.EncodeToString(encrypted[current:current+16])]))
		if previousBlockHex != "" {
			previousBlockHex = previousBlockHex[2:]
		} else {
			current += 16
			previousBlockHex = "414141414141414141414141414141"
		}
	}
	return result
}

// I separated the decryption between the first block and the rest, this is not really DRY but it works
// We do very similar things in the first block and in the rest, the only difference is that the first block
// just uses the 41s as padding, while the rest also uses the previous block, which makes it hard to write
// in one neat function.

func ecbOracle(key []byte) string {
	b1 := decryptFirstBlock(key)
	b1hex := hex.EncodeToString(b1)
	ret := string(b1)
	text := make([]byte, 0)

	padding := "414141414141414141414141414141"
	for i := 0; i < len(hexUnk()); i++ {
		m := makeByteMap(key, b1hex[2:])
		toEncrypt, _ := hex.DecodeString(padding + hexUnk())

		encrypted := ecbEncrypt(key, toEncrypt)
		text = append(text, byte(m[hex.EncodeToString(encrypted[16*(i/16)+16:16*(i/16)+32])]))
		next := fmt.Sprintf("%x", m[hex.EncodeToString(encrypted[16*(i/16)+16:16*(i/16)+32])])
		if len(next) == 0 {
			next = "00"
		} else if len(next) == 1 {
			next = "0" + next
		}
		b1hex += next
		b1hex = b1hex[2:]

		if padding != "" {
			padding = padding[2:]
		} else {
			padding = "414141414141414141414141414141"
		}
	}
	ret = ret + string(text)[1:]
	return ret
}

func main() {
	key := randomKey()
	fmt.Println(ecbOracle(key))
}

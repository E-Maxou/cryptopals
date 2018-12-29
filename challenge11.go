package main

import (
	"crypto/aes"
	"crypto/cipher"
	hex "encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func randomKey() []byte {
	key := make([]byte, 16)
	r.Read(key)
	return key
}

func getIV() []byte {
	return make([]byte, 16)
}

func cbcEncrypt(key []byte, data []byte) []byte {
	block, _ := aes.NewCipher(key)
	iv := getIV()

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(data, data)
	return data
}

func ecbEncrypt(key []byte, data []byte) []byte {
	block, _ := aes.NewCipher(key)
	for i := 0; i < len(data)-15; i += 16 {
		block.Encrypt(data[i:i+16], data[i:i+16])
	}
	return data
}

func encryptRandom(hexString string) []byte {
	data, _ := hex.DecodeString(hexString)
	padding := make([]byte, 16-len(data)%16)
	data = append(data, padding...)

	if r.Int()%2 == 0 {
		return ecbEncrypt(randomKey(), data)
	} else {
		return ecbEncrypt(randomKey(), data)
	}
}

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func main() {
	fmt.Println(encryptRandom("414141414114"))
}

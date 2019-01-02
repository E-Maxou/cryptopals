package main

import (
	"crypto/aes"
	"crypto/cipher"
	hex "encoding/hex"
	"math/rand"
	"reflect"
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
	for i := 0; i < len(data)%16; i++ {
		data = append(data, 0x04)
	}
	block, _ := aes.NewCipher(key)
	for i := 0; i+15 < len(data); i += 16 {
		block.Encrypt(data[i:i+16], data[i:i+16])
	}
	return data
}

func encryptRandom(hexString string) []byte {
	data, _ := hex.DecodeString(hexString)
	padding := make([]byte, 16-len(data)%16)
	for i := range padding {
		padding[i] = 0x04
	}
	data = append(data, padding...)

	if r.Int()%2 == 0 {
		return cbcEncrypt(randomKey(), data)
	} else {
		return ecbEncrypt(randomKey(), data)
	}
}

func bytesToBlockArray(cypher []byte) [][]byte {
	ret := make([][]byte, 0)
	for i := 0; i < len(cypher)-15; i += 16 {
		ret = append(ret, cypher[i:i+16])
	}
	return ret
}

func countDuplicates(elem [][]byte) int {
	tmp := elem
	ret := 0
	for i := range tmp {
		for j := range elem {
			if i != j && reflect.DeepEqual(tmp[i], elem[j]) {
				ret++
			}
		}
	}
	return ret
}

func oracleMode(cypher []byte) string {
	blocks := bytesToBlockArray(cypher)
	if countDuplicates(blocks) != 0 {
		return "ecb"
	} else {
		return "cbc"
	}
}

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

/*
func main() {
	for i := 0; i < 10; i++ {
		encrypted := encryptRandom("414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141")
		fmt.Println(oracleMode(encrypted))
	}
}
*/

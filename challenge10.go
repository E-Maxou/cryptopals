package main

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	hex "encoding/hex"
	"fmt"
	"log"
	"os"
)

func xor_equal_len(hex_s1 string, hex_s2 string) string {
	bytes_s1, _ := hex.DecodeString(hex_s1)
	bytes_s2, _ := hex.DecodeString(hex_s2)
	bytes_xor := make([]byte, len(bytes_s1))
	fmt.Println(bytes_s1)
	for i := 0; i < len(bytes_s1); i++ {
		bytes_xor[i] = bytes_s1[i] ^ bytes_s2[i]
	}
	return hex.EncodeToString(bytes_xor)
}

func ReadFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	bytes := make([]byte, 0, 1024)

	for scan.Scan() {
		input := []byte(scan.Text())
		bytes = append(bytes, input...)
	}
	return bytes
}

func getIV() []byte {
	return make([]byte, 16)
}

func main() {
	key := []byte("YELLOW SUBMARINE")
	file := ReadFile("10.txt")
	bytesIn, _ := base64.StdEncoding.DecodeString(string(file))
	bprec := getIV()

	AES, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	bytesOut := make([]byte, len(bytesIn))
	copy(bytesOut, bytesIn)
	dec := make([]byte, 0)
	for i := 0; i < len(bytesIn); i += 16 {
		AES.Decrypt(bytesOut[i:i+16], bytesIn[i:i+16])

		out := make([]byte, 16)
		for j := i; j < i+16; j++ {
			out[j-i] = bytesOut[j] ^ bprec[j-i]
		}
		bprec = bytesIn[i : i+16]
		dec = append(dec, out...)
	}
	fmt.Println(string(dec))
}

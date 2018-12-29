package main

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

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

func main() {
	key := []byte("YELLOW SUBMARINE")

	encoded := ReadFile("7.txt")
	bytes, _ := base64.StdEncoding.DecodeString(string(encoded))

	fmt.Println(len(bytes))
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// dst := make([]byte, 16)
	for i := 0; i < len(bytes); i += 16 {
		block.Decrypt(bytes[i:i+16], bytes[i:i+16])
	}
	fmt.Println(string(bytes))
}

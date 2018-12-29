package main

import "fmt"

func pkcs7(str string, length int) []byte {
	bytes := []byte(str)

	pkcs7 := []byte{0x04}
	fmt.Println(length)
	limit := length - len(bytes)
	for i := 0; i < limit; i++ {
		bytes = append(bytes, pkcs7...)
	}
	return bytes
}

func main() {
	a := pkcs7("YELLOW SUBMARINE", 20)
	fmt.Println("Challenge 9 : ", a, string(a))
}

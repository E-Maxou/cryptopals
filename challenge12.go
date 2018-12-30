package main

// compiles with go run challene12.go challenge11.go

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		encrypted := encryptRandom("414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141")
		fmt.Println(oracleMode(encrypted))
	}
}

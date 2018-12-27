package main

import "fmt"

import hex "encoding/hex"
import "strings"

func xor_equal_len(hex_s1 string, hex_s2 string) string {
	bytes_s1, _ := hex.DecodeString(hex_s1)
	bytes_s2, _ := hex.DecodeString(hex_s2)
	bytes_xor := make([]byte, len(bytes_s1))
	for i := 0; i < len(bytes_s1); i++ {
		bytes_xor[i] = bytes_s1[i] ^ bytes_s2[i]
	}
	return hex.EncodeToString(bytes_xor)
}

func repeatingXorMakeKey(key string, length int) string {
	ret := strings.Repeat(key, length/len(key)+1)
	return ret[0:length]
}

func repeatingKeyXor(clear string, key string) string {
	entryTab := strings.Split(clear, "\n")
	output := ""
	for i := 0; i < len(entryTab); i++ {
		repeatedKey := repeatingXorMakeKey(key, len(entryTab[i]))
		hexkey := fmt.Sprintf("%x", repeatedKey)
		hexclear := fmt.Sprintf("%x", entryTab[i])
		output += xor_equal_len(hexkey, hexclear)
		output += "\n"
	}
	return output
}

func main() {
	chall5 := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	fmt.Print("challenge 5 : ", repeatingKeyXor(chall5, "ICE"))
}

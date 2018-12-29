package main

import "fmt"

import hex "encoding/hex"
import "strings"

func decoder(str string) string {
	key := ""
	best := ""
	var max float32 = 0
	for i := 0x00; i < 0xFF; i++ {
		key = fmt.Sprintf("%c", i)
		decodedHex := xorStrToChar(str, key)
		decodedAscii := hexToAscii(decodedHex)
		if max < rateString(decodedAscii) {
			max = rateString(decodedAscii)
			best = decodedAscii
		}
	}
	if max > 1.5 {
		return best
	} else {
		return ""
	}
}

func get_standard_frequency(letter rune) float32 {
	CHARACTER_FREQ := map[rune]float32{
		'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835, 'e': 0.1041442, 'f': 0.0197881, 'g': 0.0158610,
		'h': 0.0492888, 'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490, 'm': 0.0202124, 'n': 0.0564513,
		'o': 0.0596302, 'p': 0.0137645, 'q': 0.0008606, 'r': 0.0497563, 's': 0.0515760, 't': 0.0729357, 'u': 0.0225134,
		'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692, 'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182,
	}
	return CHARACTER_FREQ[letter]
}

func rateString(str string) float32 {
	var t float32 = 0
	for _, char := range str {
		t += get_standard_frequency(char)
	}
	return t
}

func xorStrToChar(str string, character string) string {
	repeatedChar := strings.Repeat(character, len(str)/2)
	key := fmt.Sprintf("%x", repeatedChar)
	return xor_equal_len(str, key)
}

func hexToAscii(str string) string {
	bs, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func xor_equal_len(hex_s1 string, hex_s2 string) string {
	bytes_s1, _ := hex.DecodeString(hex_s1)
	bytes_s2, _ := hex.DecodeString(hex_s2)
	bytes_xor := make([]byte, len(bytes_s1))
	for i := 0; i < len(bytes_s1); i++ {
		bytes_xor[i] = bytes_s1[i] ^ bytes_s2[i]
	}
	return hex.EncodeToString(bytes_xor)
}

func main() {
	fmt.Println("challenge 3 : ", decoder("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
}

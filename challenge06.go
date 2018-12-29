package main

import "fmt"
import hex "encoding/hex"
import "strings"
import b64 "encoding/base64"
import "os"
import "bufio"

//import "reflect"

func hexToBinary(hexnum string) string {
	if hexnum == "0" {
		return "0000"
	} else if hexnum == "1" {
		return "0001"
	} else if hexnum == "2" {
		return "0010"
	} else if hexnum == "3" {
		return "0011"
	} else if hexnum == "4" {
		return "0100"
	} else if hexnum == "5" {
		return "0101"
	} else if hexnum == "6" {
		return "0110"
	} else if hexnum == "7" {
		return "0111"
	} else if hexnum == "8" {
		return "1000"
	} else if hexnum == "9" {
		return "1001"
	} else if hexnum == "a" || hexnum == "A" {
		return "1010"
	} else if hexnum == "b" || hexnum == "B" {
		return "1011"
	} else if hexnum == "c" || hexnum == "C" {
		return "1100"
	} else if hexnum == "d" || hexnum == "D" {
		return "1101"
	} else if hexnum == "e" || hexnum == "E" {
		return "1110"
	} else {
		return "1111"
	}
}

// THIS FUNCTION TAKES HEX STRING REPRESENTATIONs OF THE TWO STRINGS AS ARGUMENTS
func hamming(s1 string, s2 string) int {
	//hex1 := fmt.Sprintf("%x", s1)
	//hex2 := fmt.Sprintf("%x", s2)

	bytes_s1, _ := hex.DecodeString(s1)
	bytes_s2, _ := hex.DecodeString(s2)

	bytes_str1 := fmt.Sprintf("%x", bytes_s1)
	bytes_str1 = strings.Join(strings.Fields(bytes_str1), "")

	bytes_str2 := fmt.Sprintf("%x", bytes_s2)
	bytes_str2 = strings.Join(strings.Fields(bytes_str2), "")

	bin_str1 := ""
	bin_str2 := ""

	for _, char := range bytes_str1 {
		bin_str1 += hexToBinary(fmt.Sprintf("%c", char))
	}
	for _, char := range bytes_str2 {
		bin_str2 += hexToBinary(fmt.Sprintf("%c", char))
	}
	distance := 0

	for i := 0; i < len(bin_str1); i++ {
		if bin_str1[i] != bin_str2[i] {
			distance++
		}
	}
	return distance
}

func decodeb64(str string) []byte {
	data, _ := b64.StdEncoding.DecodeString(str)
	return data
}

func fileAsOneByteArray() []byte {
	var output []byte
	f, _ := os.Open("6.txt")
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		txt := fs.Text()
		output = append(output, decodeb64(txt)...)
	}
	return output
}

func lineAsOneByteArray() []byte {
	var output []byte
	f, _ := os.Open("6.txt")
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		txt := fs.Text()
		output = append(output, decodeb64(txt)...)
		break
	}
	return output
}

func distanceBetweenChunks(keysize int, data string) float64 {
	// we are using hex string here as input : one ascii char is two hex digits long
	out := 0.0
	totalNumberOfChunks := 0.0
	for i := 0; i < len(data)/(5*keysize); i++ {
		offset := i * 4 * keysize
		firstChunk := data[offset : offset+keysize*2]
		secondChunk := data[offset+keysize*2 : offset+4*keysize]
		res := float64(hamming(firstChunk, secondChunk))
		res /= float64(keysize)
		out += res
		totalNumberOfChunks = float64(i)
	}
	return out / totalNumberOfChunks
}

func optimalHammingDistance(data string) int {
	sizeOfMin := 0
	scoreOfMin := 1000.0
	for i := 2; i < 41; i++ {
		if scoreOfMin > distanceBetweenChunks(i, data) {
			sizeOfMin = i
			scoreOfMin = distanceBetweenChunks(i, data)
		}
	}
	return sizeOfMin
}

func cutInChunks(size int, data string) []string {
	//we are still using hex digits, thus 1 character is of length two
	var out []string
	for i := 0; i < len(data); i += size * 2 {
		if i+size*2 > len(data) {
			break
		}
		out = append(out, data[i:i+(2*size)])
	}
	return out
}

func transposeChunks(chunks []string) []string {
	var out []string
	for j := 0; j < len(chunks[0])-1; j += 2 {
		column := ""
		for i := 0; i < len(chunks); i++ {
			column += chunks[i][j : j+2]
		}
		out = append(out, column)
	}
	return out
}

/*

/*****************************************************************************************
******************************************************************************************
********************** CODE SNIPPETS FROM PREVIOUS CHALLENGES ****************************
******************************************************************************************
*****************************************************************************************/

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

func get_standard_frequency(letter rune) float32 {
	CHARACTER_FREQ := map[rune]float32{
		'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835, 'e': 0.1041442, 'f': 0.0197881, 'g': 0.0158610,
		'h': 0.0492888, 'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490, 'm': 0.0202124, 'n': 0.0564513,
		'o': 0.0596302, 'p': 0.0137645, 'q': 0.0008606, 'r': 0.0497563, 's': 0.0515760, 't': 0.0729357, 'u': 0.0225134,
		'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692, 'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182,
	}
	return CHARACTER_FREQ[letter]
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

func decoder(str string) string {
	key := ""
	keyFound := ""
	//best := ""
	var max float32 = 0
	for i := 0x00; i < 0xFF; i++ {
		key = fmt.Sprintf("%c", i)
		decodedHex := xorStrToChar(str, key)
		decodedAscii := hexToAscii(decodedHex)
		if max < rateString(decodedAscii) {
			//fmt.Println(max)
			max = rateString(decodedAscii)
			//	best = decodedAscii
			keyFound = key
		}
	}
	if max > 1.5 {
		return keyFound
	} else {
		return ""
	}
}

func main() {
	fileAsHex := hex.EncodeToString(fileAsOneByteArray())
	fmt.Println("The length of the keysize is likely :", optimalHammingDistance(fileAsHex))

	chunks := cutInChunks(29, fileAsHex)
	T := transposeChunks(chunks)
	for i := 0; i < len(T); i++ {
		fmt.Println(i, "th char :", decoder(T[i]))
	}
}

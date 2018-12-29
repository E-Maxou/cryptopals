package main

import "fmt"
import "os"
import "bufio"
import hex "encoding/hex"
import "reflect"

func countDuplicates(elem [10][]byte) int {
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

func strToBlockArray(cypher string) [10][]byte {
	hexa, _ := hex.DecodeString(cypher)
	var ret [10][]byte
	for i := 0; i < len(hexa)-15; i += 16 {
		ret[i/16] = hexa[i : i+16]
	}
	return ret
}

func main() {
	f, _ := os.Open("8.txt")
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		txt := fs.Text()
		e := strToBlockArray(txt)
		if countDuplicates(e) != 0 {
			fmt.Println("Challenge 8,the encoded block is :", txt)
		}
	}
}

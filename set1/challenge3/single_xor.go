package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	decoded_bytes, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	fmt.Println(decoded_bytes)

	n := len(decoded_bytes)

	for k := range 256 {
		xored := []byte{}
		for i := range n {
			x := decoded_bytes[i] ^ byte(k)
			xored = append(xored, x)
		}

		encoded_str := string(xored)
		fmt.Println("K: ", k, "> ", encoded_str)
	}
}

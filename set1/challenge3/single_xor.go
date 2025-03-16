package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var freq string = "etaoin shrdlu"

	decoded_bytes, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	fmt.Println(decoded_bytes)

	n := len(decoded_bytes)

	total_score := 0
	var candidate string
	for k := range 256 {
		xored := []byte{}
		for i := range n {
			x := decoded_bytes[i] ^ byte(k)
			xored = append(xored, x)
		}

		encoded_str := string(xored)
		score := 0
		for _, char := range freq {
			c := strings.Count(encoded_str, string(char))
			score += c
		}

		if score > total_score {
			total_score = score
			candidate = encoded_str
		}

	}

	fmt.Println(candidate)
}

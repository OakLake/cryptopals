package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	var freq string = "etaoin shrdlu"

	dat, err := os.ReadFile("./challenge4/four.txt")
	checkError(err)

	lines := strings.Split(string(dat), "\n")

	total_score := 0
	var candidate string
	for _, line := range lines {

		decoded_bytes, _ := hex.DecodeString(line)

		n := len(decoded_bytes)

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

	}
	fmt.Println(candidate)
}

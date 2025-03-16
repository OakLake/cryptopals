package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	first := "1c0111001f010100061a024b53535009181c"
	second := "686974207468652062756c6c277320657965"

	first_bytes, err := hex.DecodeString(first)

	if err != nil {
		return
	}

	second_bytes, err := hex.DecodeString(second)

	if err != nil {
		return
	}

	n := len(first_bytes)
	xor_result := []byte{}
	for i := range n {
		x := first_bytes[i] ^ second_bytes[i]
		xor_result = append(xor_result, x)
	}

	fmt.Println(first_bytes)
	fmt.Println(second_bytes)
	fmt.Println(xor_result)

	hex_result := hex.EncodeToString(xor_result)

	fmt.Println("Result > ", hex_result)
}

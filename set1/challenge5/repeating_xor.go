package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	text := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")

	locked := []byte{}
	for i := range text {
		x := text[i] ^ key[i%3]
		locked = append(locked, x)
	}

	locked_text := hex.EncodeToString(locked)
	fmt.Println(locked_text)

}

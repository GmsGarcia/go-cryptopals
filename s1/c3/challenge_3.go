package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	const ciphertext = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	const keys = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hex1, err1 := hex.DecodeString(ciphertext)

	if err1 != nil {
		panic(err1)
	}

	n := len(hex1)

	for j, c := range keys {
		xored := make([]byte, n)
		for i := 0; i < n; i++ {
			xored[i] = hex1[i] ^ byte(c)
		}
		fmt.Printf("\n%2d %c - %s", j, c, string(xored))

		// Letter X - 'Cooking MC's like a pound of bacon.'
	}
}

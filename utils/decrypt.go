package utils

import (
	"encoding/hex"
	"fmt"
)

func DecryptSingleCharacterXOR(ciphertext string) {
	fmt.Printf("\nEncoded string: %s", ciphertext)

	hex, err := hex.DecodeString(ciphertext)
	HandleErrorAgressively(err)

	n := len(hex)

	for c := 0; c <= 255; c++ {
		xored := make([]byte, n)
		for i := 0; i < n; i++ {
			xored[i] = hex[i] ^ byte(c)
		}
		fmt.Printf("\n%2d %c - %s", c, byte(c), string(xored))

		// Letter X - 'Cooking MC's like a pound of bacon.'
	}
}

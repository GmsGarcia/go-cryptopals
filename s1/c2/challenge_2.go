package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hex1, err1 := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	hex2, err2 := hex.DecodeString("686974207468652062756c6c277320657965")

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	n := len(hex1)
	xored := make([]byte, n)

	for i := 0; i < n; i++ {
		xored[i] = hex1[i] ^ hex2[i]
	}

	result := hex.EncodeToString(xored)

	fmt.Print(result)
}

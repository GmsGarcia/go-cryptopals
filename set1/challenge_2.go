package set1

import (
	"encoding/hex"
	"fmt"
	"go-learn/cryptopals/utils"
)

func C2() {
	const data1 = "1c0111001f010100061a024b53535009181c"
	const data2 = "686974207468652062756c6c277320657965"

	fmt.Printf("\nFirst string: %v", data1)
	fmt.Printf("\nSecond string: %v", data2)

	hex1, err1 := hex.DecodeString(data1)
	utils.HandleErrorAgressively(err1)

	hex2, err2 := hex.DecodeString(data2)
	utils.HandleErrorAgressively(err2)

	n := len(hex1)
	xored := make([]byte, n)

	for i := 0; i < n; i++ {
		xored[i] = hex1[i] ^ hex2[i]
	}

	result := hex.EncodeToString(xored)
	fmt.Printf("\nXOR'd string: %s", result)
}

package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"go-learn/cryptopals/utils"
)

func C1() {
	const data = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Printf("\nHex string: %s", data)

	bytes, err := hex.DecodeString(data)
	utils.HandleErrorAgressively(err)

	result := base64.StdEncoding.EncodeToString(bytes)
	fmt.Printf("\nBase64 string: %s", result)
}

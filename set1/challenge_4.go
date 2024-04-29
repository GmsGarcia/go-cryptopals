package set1

import (
	"encoding/hex"
	"fmt"
	"go-learn/cryptopals/utils"
	"io"
	"net/http"
	"strings"
)

func C4() {
	arr := getStringsFromWeb()
	crackXORCipher(arr[1])
}

func crackXORCipher(cipherString string) string {
	cipherBytes, err1 := hex.DecodeString(cipherString)
	utils.HandleErrorAgressively(err1)

	var bestString string

	for k := 0; k < 256; k++ {
		plainText := make([]byte, len(cipherBytes))

		for i := 0; i < len(cipherBytes); i++ {
			plainText[i] = cipherBytes[i] ^ byte(k)
		}
		fmt.Printf("%c - %s", k, plainText)
	}

	return bestString
}

func getStringsFromWeb() []string {
	res, err1 := http.Get("https://cryptopals.com/static/challenge-data/4.txt")
	utils.HandleErrorAgressively(err1)

	defer res.Body.Close()

	body, err2 := io.ReadAll(res.Body)
	utils.HandleErrorAgressively(err2)

	return strings.Split(string(body), "\n")
}

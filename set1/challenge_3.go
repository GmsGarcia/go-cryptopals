package set1

import (
	"go-learn/cryptopals/utils"
)

func C3() {
	const ciphertext = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	utils.DecryptSingleCharacterXOR(ciphertext)
}

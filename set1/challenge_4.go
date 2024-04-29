package set1

import (
	"go-learn/cryptopals/utils"
)

func C4() {
	arr := utils.GetStringsFromWeb()
	utils.DecryptSingleCharacterXOR(arr[1])
}

package main

import (
	cmenu "go-learn/cryptopals/menu"
	"go-learn/cryptopals/set1"
	"go-learn/cryptopals/utils"
)

func main() {
	prompt := cmenu.GetPrompt()

	i, _, err := prompt.Run()
	utils.HandleErrorPassively(err)

	switch i {
	case 0:
		set1.C1()
	case 1:
		set1.C2()
	case 2:
		set1.C3()
	case 3:
		set1.C4()
	}
}

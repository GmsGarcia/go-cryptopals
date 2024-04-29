package utils

import "fmt"

func HandleErrorAgressively(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleErrorPassively(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

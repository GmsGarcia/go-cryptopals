package utils

import (
	"io"
	"net/http"
	"strings"
)

func GetStringsFromWeb() []string {
	res, err1 := http.Get("https://cryptopals.com/static/challenge-data/4.txt")
	HandleErrorAgressively(err1)

	defer res.Body.Close()

	body, err2 := io.ReadAll(res.Body)
	HandleErrorAgressively(err2)

	return strings.Split(string(body), "\n")
}

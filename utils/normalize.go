package utils

import (
	"log"
	"regexp"
	"strings"
)

func NormalizeText(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal("normalization regex failed")
	}
	return strings.ToUpper(reg.ReplaceAllString(s, ""))
}

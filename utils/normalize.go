package utils

import (
	"log"
	"regexp"
	"strings"
)

func NormalizeText(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}

	return strings.ToUpper(strings.TrimSpace(reg.ReplaceAllString(s, "")))
}

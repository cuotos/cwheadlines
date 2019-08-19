package morse

import "strings"

const (
	Dot          = 1
	Dash         = 3
	InterCharGap = 1
	CharGap      = 3
	WordGap      = 7

	wpm = 1 // wpm = "PARIS"s/min

	wordSeperator   = "/"
	letterSeperator = " "
)

var runeMorseMap = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",
	' ': "/",
}

func RuneToMorse(input rune) string {
	return runeMorseMap[input]
}

func Encode(input string) string {
	res := ""

	for _, c := range input {
		if string(c) == " " {
			res += wordSeperator + letterSeperator
		} else {
			res += runeMorseMap[c] + letterSeperator
		}
	}

	return strings.TrimSpace(res)
}

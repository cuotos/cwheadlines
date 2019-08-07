package morse

import (
	"bytes"
	"strings"
)
var runeMorseMap = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': ".-..",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
	' ': "/",
}

var morseMap = map[string]string{
	"a": `᛫ -`,
	"b": "- ᛫ ᛫ ᛫",
	"c": "- ᛫ - ᛫",
	"d": "- ᛫ ᛫",
	"e": "᛫",
	"f": "᛫ ᛫ - ᛫",
	"g": "- - ᛫",
	"h": "᛫ ᛫ ᛫ ᛫",
	"i": "᛫ ᛫",
	"j": "᛫ - - -",
	"k": "- ᛫ -",
	"l": "᛫ - ᛫ ᛫",
	"m": "- -",
	"n": "- ᛫",
	"o": "- - -",
	"p": "᛫- - ᛫",
	"q": "- - ᛫ -",
	"r": "᛫ - ᛫",
	"s": "᛫ ᛫ ᛫",
	"t": "-",
	"u": "᛫ ᛫ -",
	"v": "᛫ ᛫ ᛫ -",
	"w": "᛫ - -",
	"x": "- ᛫ ᛫ -",
	"y": "- ᛫ - -",
	"z": "- - ᛫ ᛫",
	"0": "- - - - -",
	"1": "᛫ - - - -",
	"2": "᛫ ᛫ - - -",
	"3": "᛫ ᛫ ᛫ - -",
	"4": "᛫ ᛫ ᛫ ᛫ -",
	"5": "᛫ ᛫ ᛫ ᛫ ᛫",
	"6": "- ᛫ ᛫ ᛫ ᛫",
	"7": "- - ᛫ ᛫ ᛫",
	"8": "- - - ᛫ ᛫",
	"9": "- - - - ᛫",
	" ": " ",
}

func RuneToMorse(input rune) string {
	return runeMorseMap[input]
}

func AsMorse(input string) string {
	input = strings.ToLower(input)
	var buffer bytes.Buffer

	for _, c := range input {
		key := string(c)
		_, exists := morseMap[key]
		if exists {
			buffer.WriteString(morseMap[key])
			buffer.WriteString(" / ")
		}
	}

	return buffer.String()
}
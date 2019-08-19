package morse

import (
	"log"
	"strings"
	"unicode"
)

var (
	Dot          = Key{1, true, "."}
	Dash         = Key{3, true, "-"}
	InterCharGap = Key{1, false, ""}
	CharGap      = Key{3, false, " "}
	WordGap      = Key{7, false, " / "}
)

// TODO extract this to a string so it cant be used directly
// for reference only, use runeToMorseChars map
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
}

var runeToKeys = map[rune]Keys{}

// populate runeToMorse map
func init() {
	for k, v := range runeMorseMap {
		keys := []Key{}

		for i, s := range v {
			if i > 0 {
				keys = append(keys, InterCharGap)
			}
			if s == '.' {
				keys = append(keys, Dot)
			} else if s == '-' {
				keys = append(keys, Dash)
			} else {
				log.Fatalf("found %s in morse for %v", string(s), v)
			}
		}

		runeToKeys[k] = keys
	}
}

type Keys []Key

func (k Keys) String() string {
	res := strings.Builder{}

	for _, k := range k {
		res.WriteString(k.Symb)
	}

	return res.String()
}

type Key struct {
	Duration int
	On       bool
	Symb     string
}

func StringToKeys(input string) Keys {

	keys := Keys{}
	afterChar := false

	for _, c := range input {

		if unicode.IsSpace(c) {
			keys = append(keys, WordGap)
			afterChar = false
			continue
		}

		if afterChar {
			keys = append(keys, CharGap)
			afterChar = false
		}
		switch c {
		case ' ':
			keys = append(keys, WordGap)
		default:
			keys = append(keys, runeToKeys[c]...)
		}
		afterChar = true
	}

	return keys
}

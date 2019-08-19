package problem

import (
	"cwheadlines/morse"
	"cwheadlines/utils"
)

func GenerateProblem(input string) Problem {

	input = utils.NormalizeText(input)

	p := Problem{}

	p.Solution = input

	p.Morse = morse.StringToKeys(input)

	return p
}

package problem

import (
	"cwheadlines/morse"
	"cwheadlines/utils"
)

func GenerateProblem(input string) Problem {

	input = utils.NormalizeText(input)

	p := Problem{}

	for _, r := range input {
		c := &Character{}
		c.solution = r
		c.morse = morse.RuneToMorse(r)

		p = append(p, c)
	}

	return p
}

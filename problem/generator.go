package problem

import (
	"cwheadlines/morse"
)

func GenerateProblem(input string) Problem {

	p := Problem{}

	for _, r := range input {
		c := &Character{}
		c.solution = r
		c.morse = morse.RuneToMorse(r)

		p = append(p, c)
	}

	return p
}

package problem

import (
	"bytes"
	"cwheadlines/morse"
	"strings"
)

type Problem []*Character

func (p Problem) AsMorse() string {
	var buffer bytes.Buffer

	for _, c := range p {
		buffer.WriteString(c.morse + " ")
	}

	return strings.TrimSpace(buffer.String())
}

type Character struct {
	solution int32
	morse    string
	correct  bool
}

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

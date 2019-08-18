package problem

import (
	"bytes"
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

func (p Problem) AsString() string {
	var buffer bytes.Buffer

	for _, c := range p {
		buffer.WriteString(string(c.solution))
	}

	return buffer.String()
}

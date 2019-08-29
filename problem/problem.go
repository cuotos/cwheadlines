package problem

import (
	"github.com/cuotos/cwheadlines/morse"
)

type Problem struct {
	Solution string
	Morse    morse.Keys
}

func (p Problem) AsMorse() string {
	return p.Morse.String()
}

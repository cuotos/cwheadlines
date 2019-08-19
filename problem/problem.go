package problem

import (
	"cwheadlines/morse"
)

type Problem struct {
	Solution string
	Morse    morse.Keys
}

func (p Problem) AsMorse() string {
	return p.Morse.String()
}

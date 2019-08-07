package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblemGenerator(t *testing.T) {
	tcs := []struct{
		name string
		input string
		expected Problem
	}{
		{
			"single character",
			"a",
			Problem{
				&Character{'a', ".-", false}},
		},
		{
			"three letters",
			"abc",
			Problem{
				&Character{'a', ".-", false},
				&Character{'b', "-...", false},
				&Character{'c', "-.-.", false},
			},
		},
		{
			"test",
			"test",
			Problem{
				&Character{'t', "-", false},
				&Character{'e', ".", false},
				&Character{'s', "...", false},
				&Character{'t', "-", false},
			},
		},
		{
			"multi word",
			"abc def",
			Problem{
				&Character{'a', ".-", false},
				&Character{'b', "-...", false},
				&Character{'c', "-.-.", false},
				&Character{' ', "/", false},
				&Character{'d', "-..", false},
				&Character{'e', ".", false},
				&Character{'f', ".-..", false},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			actual := GenerateProblem(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPrintProblemAsString(t *testing.T) {
	tcs := []struct{
		name string
		input string
		expected string
	}{
		{
			"single word",
			"test",
			"- . ... -",
		},
		{
			"simple sentence",
			"this is a test",
			"- .... .. ... / .. ... / .- / - . ... -",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := GenerateProblem(tc.input)
			actual := p.AsMorse()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblemGenerator(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected Problem
	}{
		{
			"single character",
			"a",
			Problem{
				&Character{'A', ".-", false}},
		},
		{
			"three letters",
			"abc",
			Problem{
				&Character{'A', ".-", false},
				&Character{'B', "-...", false},
				&Character{'C', "-.-.", false},
			},
		},
		{
			"test",
			"test",
			Problem{
				&Character{'T', "-", false},
				&Character{'E', ".", false},
				&Character{'S', "...", false},
				&Character{'T', "-", false},
			},
		},
		{
			"multi word",
			"abc def",
			Problem{
				&Character{'A', ".-", false},
				&Character{'B', "-...", false},
				&Character{'C', "-.-.", false},
				&Character{' ', "/", false},
				&Character{'D', "-..", false},
				&Character{'E', ".", false},
				&Character{'F', "..-.", false},
			},
		},
		{
			"alphabet",
			"abcdefghijklmnopqrstuvwxyz",
			Problem{
				&Character{'A', ".-", false},
				&Character{'B', "-...", false},
				&Character{'C', "-.-.", false},
				&Character{'D', "-..", false},
				&Character{'E', ".", false},
				&Character{'F', "..-.", false},
				&Character{'G', "--.", false},
				&Character{'H', "....", false},
				&Character{'I', "..", false},
				&Character{'J', ".---", false},
				&Character{'K', "-.-", false},
				&Character{'L', ".-..", false},
				&Character{'M', "--", false},
				&Character{'N', "-.", false},
				&Character{'O', "---", false},
				&Character{'P', ".--.", false},
				&Character{'Q', "--.-", false},
				&Character{'R', ".-.", false},
				&Character{'S', "...", false},
				&Character{'T', "-", false},
				&Character{'U', "..-", false},
				&Character{'V', "...-", false},
				&Character{'W', ".--", false},
				&Character{'X', "-..-", false},
				&Character{'Y', "-.--", false},
				&Character{'Z', "--..", false},
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

func TestPrintProblemAsMorse(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
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

func TestPrintProblemAsString(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"simple word",
			"TheInput",
			"THEINPUT",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			p := GenerateProblem(tc.input)
			actual := p.AsString()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

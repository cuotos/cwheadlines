package problem

import (
	. "github.com/cuotos/cwheadlines/morse"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	A   = []Key{Dot, InterCharGap, Dash}
	B   = []Key{Dash, InterCharGap, Dot, InterCharGap, Dot, InterCharGap, Dot}
	C   = []Key{Dash, InterCharGap, Dot, InterCharGap, Dash, InterCharGap, Dot}
	D   = []Key{Dash, InterCharGap, Dot, InterCharGap, Dot}
	E   = []Key{Dot}
	F   = []Key{Dot, InterCharGap, Dot, InterCharGap, Dash, InterCharGap, Dot}
	S   = []Key{Dot, InterCharGap, Dot, InterCharGap, Dot}
	T   = []Key{Dash}
	Gap = []Key{CharGap}
)

func JoinSlices(slices ...[]Key) []Key {
	kys := []Key{}

	for _, s := range slices {
		for _, k := range s {
			kys = append(kys, k)
		}
	}

	return kys
}

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
				"A",
				[]Key{Dot, InterCharGap, Dash},
			},
		},
		{
			"three letters",
			"abc",
			Problem{
				"ABC",
				JoinSlices(A, Gap, B, Gap, C),
			},
		},
		{
			"test",
			"test",
			Problem{
				"TEST",
				JoinSlices(T, Gap, E, Gap, S, Gap, T),
			},
		},
		{
			"multi word",
			"abc def",
			Problem{
				"ABC DEF",
				JoinSlices(A, Gap, B, Gap, C, []Key{WordGap}, D, Gap, E, Gap, F),
			},
		},
		// TODO: reenable full alphabet test
		//{
		//	"alphabet",
		//	"abcdefghijklmnopqrstuvwxyz",
		//	Problem{
		//		&Character{'A', ".-", false},
		//		&Character{'B', "-...", false},
		//		&Character{'C', "-.-.", false},
		//		&Character{'D', "-..", false},
		//		&Character{'E', ".", false},
		//		&Character{'F', "..-.", false},
		//		&Character{'G', "--.", false},
		//		&Character{'H', "....", false},
		//		&Character{'I', "..", false},
		//		&Character{'J', ".---", false},
		//		&Character{'K', "-.-", false},
		//		&Character{'L', ".-..", false},
		//		&Character{'M', "--", false},
		//		&Character{'N', "-.", false},
		//		&Character{'O', "---", false},
		//		&Character{'P', ".--.", false},
		//		&Character{'Q', "--.-", false},
		//		&Character{'R', ".-.", false},
		//		&Character{'S', "...", false},
		//		&Character{'T', "-", false},
		//		&Character{'U', "..-", false},
		//		&Character{'V', "...-", false},
		//		&Character{'W', ".--", false},
		//		&Character{'X', "-..-", false},
		//		&Character{'Y', "-.--", false},
		//		&Character{'Z', "--..", false},
		//	},
		//},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			actual := GenerateProblem(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPrintProblemAsMorseString(t *testing.T) {
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

func TestSolutionCorrectlyNormalized(t *testing.T) {
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
			actual := p.Solution
			assert.Equal(t, tc.expected, actual)
		})
	}
}

package morse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	tcs := []struct {
		input    rune
		expected Keys
	}{
		{
			'A',
			Keys{Dot, InterCharGap, Dash},
		},
		{
			'B',
			Keys{Dash, InterCharGap, Dot, InterCharGap, Dot, InterCharGap, Dot},
		},
	}

	for _, tc := range tcs {
		actual := runeToKeys[tc.input]
		assert.Equal(t, tc.expected, actual)
	}
}

func TestMorseToKeys(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected Keys
	}{
		{"A",
			"A",
			Keys{Dot, InterCharGap, Dash},
		},
		{"space",
			" ",
			Keys{Key{7, false, " / "}},
		},
		{"AB",
			"AB",
			[]Key{Dot, InterCharGap, Dash, CharGap, Dash, InterCharGap, Dot, InterCharGap, Dot, InterCharGap, Dot},
		},
		{
			"A B",
			"A B",
			[]Key{Dot, InterCharGap, Dash, WordGap, Dash, InterCharGap, Dot, InterCharGap, Dot, InterCharGap, Dot},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := StringToKeys(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestKeysToString(t *testing.T) {
	tcs := []struct {
		input    Keys
		expected string
	}{
		{
			Keys{Dot, InterCharGap, Dash},
			".-",
		},
		{
			Keys{Dot, InterCharGap, Dash, CharGap, Dot, InterCharGap, Dot},
			".- ..",
		},
		{
			Keys{Dot, InterCharGap, Dash, WordGap, Dot, InterCharGap, Dot},
			".- / ..",
		},
	}

	for _, tc := range tcs {
		actual := tc.input.String()
		assert.Equal(t, tc.expected, actual)
	}
}

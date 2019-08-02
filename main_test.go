package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeText(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Simple Caps",
			"This is a Test Title that only contains capitals",
			"this is a test title that only contains capitals",
		},
		{
			"Quotes",
			`dan says "I am amazing"`,
			"dan says i am amazing",
		},
		{
			"Quots 2",
			`Plane lands on road and gets 'pulled over'`,
			"plane lands on road and gets pulled over",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := normalizeText(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

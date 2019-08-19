package utils

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
			"THIS IS A TEST TITLE THAT ONLY CONTAINS CAPITALS",
		},
		{
			"Quotes",
			`dan says "I am amazing"`,
			"DAN SAYS I AM AMAZING",
		},
		{
			"Quots 2",
			`Plane lands on road and gets 'pulled over'`,
			"PLANE LANDS ON ROAD AND GETS PULLED OVER",
		},
		// TODO test question marks etc
		// TODO test multiline input
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := NormalizeText(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

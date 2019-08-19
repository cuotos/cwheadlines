package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuessHandler(t *testing.T) {

	tcs := []struct {
		name            string
		solution        string
		input           string
		expectedCorrect bool
	}{
		{
			"simple correct answer",
			"SOLUTION",
			"solution",
			true,
		},
		{
			"simple incorrect answer",
			"SOLUTION",
			"wrong",
			false,
		},
		{
			"correctly guessed first letter",
			"SOLUTION",
			"s",
			false,
		},
		{
			"odd characters in guess",
			"SOLUTION",
			"±§~&^%",
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			correct, _ := GuessHandler(tc.solution, tc.input)
			assert.Equal(t, tc.expectedCorrect, correct)
		})
	}

}

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
			"solution",
			"solution",
			true,
		},
		{
			"simple incorrect answer",
			"solution",
			"wrong",
			false,
		},
		{
			"correctly guessed first letter",
			"solution",
			"s",
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

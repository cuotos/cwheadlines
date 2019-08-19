package morse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO test letters and characters
func TestEncode(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"simple word",
			"TEST",
			"- . ... -",
		},
		{
			"simple multi word",

			"TEST PARIS",
			"- . ... - / .--. .- .-. .. ...",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Encode(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

// TODO test decoding

package core

import (
	"testing"
)

func TestTokenizesHeaders(t *testing.T) {
	var tests = []struct {
		input string
		want  Token
	}{
		{"#Heading", NewHeaderToken("Heading", 1)},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestTokenizesHeaders(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

package core

import (
	"testing"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func TestTokenizesHeaders(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Block
	}{
		{"# Heading", blocks.NewHeaderBlock("Heading", 1)},
		{"## Heading", blocks.NewHeaderBlock("Heading", 2)},
		{"### Heading", blocks.NewHeaderBlock("Heading", 3)},
	}

	for _, test := range tests {
		if result := Tokenize(test.input); result[0] != test.want {
			t.Errorf("TestTokenizesHeaders(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

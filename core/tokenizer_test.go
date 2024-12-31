package core

import (
	"testing"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func TestBlockizesHeaders(t *testing.T) {
	var tests = []struct {
		input string
		want  blocks.Block
	}{
		{"# Heading", blocks.NewHeaderBlock("Heading", 1)},
		{"## Heading", blocks.NewHeaderBlock("Heading", 2)},
		{"### Heading", blocks.NewHeaderBlock("Heading", 3)},
	}

	for _, test := range tests {
		if result := Blockize(test.input); result[0] != test.want {
			t.Errorf("TestBlockizesHeaders(%s) = %q, want %q", test.input, result, test.want)
		}
	}
}

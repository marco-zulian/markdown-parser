package core

import (
  "testing"

  "github.com/marco-zulian/markdown-parser/blocks"
)

func TestShouldDetectCodeBlocks(t *testing.T) {
  var tests = []struct {
    input string
    want  blocks.Block
  }{
    {"    a simple\n      indented code block", blocks.NewCodeBlock("a simple\n  indented code block")}, 
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestShouldDetectCodeBlocks(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

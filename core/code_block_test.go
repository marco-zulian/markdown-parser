package core

import (
  "testing"

  "github.com/marco-zulian/markdown-parser/blocks"
)

func TestShouldDetectCodeBlocks(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"    a simple\n      indented code block", "Code: a simple\n  indented code block"}, 
  }

  for _, test := range tests {
    if result := blocks.GenerateBlock(test.input); result.String() != test.want {
      t.Errorf("TestShouldDetectCodeBlocks(%s) = %s, want %s", test.input, result, test.want)
    }
  }
}

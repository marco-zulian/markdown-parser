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
    {"mock_files/code.md", "Code: a simple\n  indented code block\n"}, 
  }

  for _, test := range tests {
    if result, _ := Tokenize(test.input); *result.GetContent() != test.want {
      t.Errorf("TestShouldDetectCodeBlocks(%s) = \n%q, want \n%q", test.input, *result.GetContent(), test.want)
    }
  }
}

func TestTabsShouldOnlyBeExpandedWhenTheyHelpDefiningStructure(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"\tfoo\tbaz\t\tbin", "Code: foo\tbaz\t\tbin"},
    {"  \tfoo\tbaz\t\tbim", "Code: foo\tbaz\t\tbim"},
  }

  for _, test := range tests {
    if result := blocks.GenerateBlock(test.input); result.String() != test.want {
      t.Errorf("TestTabsShouldOnlyBeExpandedWhenTheyHelpDefiningStructure(%s) = %s, want %s", test.input, result, test.want)
    }
  }
}

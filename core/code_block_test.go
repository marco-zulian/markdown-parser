package core

import (
  "bytes"
  "testing"
  
  "github.com/marco-zulian/markdown-parser/blocks"
)

func TestCodeBlocksShouldBePreceededByAtLeastFourSpaces(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"    a simple\n      indented code block", "Code: a simple\n  indented code block\n"},
    {"      foo\n    bar", "Code:   foo\nbar\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestCodeBlocksShouldBePreceededByAtLeastFourSpaces(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
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

func TestCodeBlockContentsAreLiteralText(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"    <a/>\n    *hi*\n\n    - one", "Code: <a/>\n*hi*\n\n- one\n"},
    {"    chunk1\n\n    chunk2\n  \n \n \n    chunk3", "Code: chunk1\n\nchunk2\n\n\n\nchunk3\n"},
    {"    chunk1\n      \n      chunk2", "Code: chunk1\n  \n  chunk2\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestCodeBlockContentsAreLiteralText(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}

func TestCodeBlockCantInterruptParagraph(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"Foo\n    bar", "Paragraph: Foo\nbar\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestCodeBlockCantInterruptParagraph(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  } 
}

func TestCodeBlockShouldBeInterruptedByAnyLineWithFewerThanFourSpaces(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"    foo\nbar", "Code: foo\nParagraph: bar\n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestCodeBlockShouldBeInterruptedByAnyLineWithFewerThanFourSpaces(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }  
}

func TestTraillingSpacesAreIncludedOnBlockContent(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"    foo  ", "Code: foo  \n"},
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestTraillingSpacesAreIncludedOnBlockContent(%s) = %s, want %s", test.input, *result.GetContent(), test.want)
    }
  }
}

func TestBlankLinesPreceedingOrFollowingAreNotIncluded(t *testing.T) {
  var tests = []struct {
    input string
    want  string
  }{
    {"\n    \n    foo\n    ", "Code: foo\n"},    
    {"\n    \n    foo\n       \n     \n", "Code: foo\n"},    
  }

  for _, test := range tests {
    buf := bytes.NewBufferString(test.input)
    if result, _ := GenerateBlockStructure(buf); *result.GetContent() != test.want {
      t.Errorf("TestBlankLinesPreceedingOrFollowingAreNotIncluded(%s) = %q, want %q", test.input, *result.GetContent(), test.want)
    }
  }
}

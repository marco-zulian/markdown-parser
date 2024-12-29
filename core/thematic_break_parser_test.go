package core

import (
  "testing"

  "github.com/marco-zulian/markdown-parser/blocks"
)

func TestParsesBreaksWithOnlyThreeCorrectTypesOfCharacter(t *testing.T) {
  var tests = []struct {
    input string
    want  blocks.Token
  }{
    {"***", blocks.NewThematicBreakBlock()},
    {"---", blocks.NewThematicBreakBlock()},
    {"___", blocks.NewThematicBreakBlock()},
    {"+++", blocks.NewParagraphBlock("+++")},
    {"===", blocks.NewParagraphBlock("===")},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestParsesBreaksWithOnlyThreeCorrectTypesOfCharacter(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

func TestAtLeastThreeCharactersShouldAppear(t *testing.T) {
  var tests = []struct {
    input string
    want  blocks.Token
  }{
    {"**", blocks.NewParagraphBlock("**")},
    {"--", blocks.NewParagraphBlock("--")},
    {"__", blocks.NewParagraphBlock("__")},
    {"_____________________________________", blocks.NewThematicBreakBlock()},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestAtLeastThreeCharactersShouldAppear(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

func TestShouldAllowSpacesAndTabsBetweenCharacters(t *testing.T) {
  var tests = []struct {
    input string
    want  blocks.Token
  }{
    {" - - -", blocks.NewThematicBreakBlock()},
    {" **  * ** * ** * **", blocks.NewThematicBreakBlock()},
    {"-   -   -    -", blocks.NewThematicBreakBlock()},
    {"- - - -    ", blocks.NewThematicBreakBlock()},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestShouldAllowSpacesAndTabsBetweenCharacters(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

func TestShouldOnlyMatchWhenOnlyTheSameSpecialCharacterAppears(t *testing.T) {
  var tests = []struct {
    input string
    want  blocks.Token
  }{
    {"_ _ _ _ a", blocks.NewParagraphBlock("_ _ _ _ a")},
    {"a------", blocks.NewParagraphBlock("a------")},
    {"---a---", blocks.NewParagraphBlock("---a---")},
    {" *-*", blocks.NewParagraphBlock(" *-*")},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestShouldOnlyMatchWhenOnlyTheSameSpecialCharacterAppears(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

func TestShouldAllowUpToThreeSpacesOfIdentation(t *testing.T) {
  var tests = []struct {
    input string
    want  blocks.Token
  }{
    {" ***", blocks.NewThematicBreakBlock()},
    {"  ---", blocks.NewThematicBreakBlock()},
    {"   ___", blocks.NewThematicBreakBlock()},
    {"    ***", blocks.NewParagraphBlock("    ***")},
  }

  for _, test := range tests {
    if result := Tokenize(test.input); result[0] != test.want {
      t.Errorf("TestShouldAllowUpToThreeSpacesOfIdentation(%s) = %q, want %q", test.input, result, test.want)
    }
  }
}

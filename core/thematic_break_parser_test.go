package core

import (
	"testing"

	"github.com/marco-zulian/markdown-parser/blocks"
)

func TestParsesBreaksWithOnlyThreeCorrectTypesOfCharacter(t *testing.T) {
	var tests = []struct {
		input string
		want  string 
	}{
		{"***", "Thematic break"},
		{"---", "Thematic break"},
		{"___", "Thematic break"}, 
    {"+++", "Paragraph: +++"},
    {"===", "Paragraph: ==="},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestParsesBreaksWithOnlyThreeCorrectTypesOfCharacter(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestAtLeastThreeCharactersShouldAppear(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
    {"**", "Paragraph: **"},
    {"--", "Paragraph: --"},
    {"__", "Paragraph: __"},
		{"_____________________________________", "Thematic break"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestAtLeastThreeCharactersShouldAppear(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestShouldAllowSpacesAndTabsBetweenCharacters(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{" - - -", "Thematic break"},
		{" **  * ** * ** * **", "Thematic break"},
		{"-   -   -    -", "Thematic break"},
		{"- - - -    ", "Thematic break"}, 
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestShouldAllowSpacesAndTabsBetweenCharacters(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestShouldOnlyMatchWhenOnlyTheSameSpecialCharacterAppears(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
      {"_ _ _ _ a", "Paragraph: _ _ _ _ a"},
      {"a------", "Paragraph: a------"},
      {"---a---", "Paragraph: ---a---"},
      {" *-*", "Paragraph:  *-*"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestShouldOnlyMatchWhenOnlyTheSameSpecialCharacterAppears(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

func TestShouldAllowUpToThreeSpacesOfIdentation(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{" ***", "Thematic break"},
		{"  ---", "Thematic break"}, 
		{"   ___", "Thematic break"}, 
    {"    ***", "Code: ***"},
	}

	for _, test := range tests {
		if result := blocks.GenerateBlock(test.input); result.String() != test.want {
			t.Errorf("TestShouldAllowUpToThreeSpacesOfIdentation(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}

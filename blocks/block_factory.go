package blocks

import (
  "regexp"
  "strings"
)

func GenerateBlock(line string) Block {

  var blockTypeRegexs = map[BlockType]*regexp.Regexp{
    Header        : regexp.MustCompile(`^( {0,3})(#{1,6})([ \t]+|$)`),
    ThematicBreak : regexp.MustCompile(`^((\*{3,})|(-{3,})|(_{3,}))$`),
    Code          : regexp.MustCompile(`^ {4,}`),
  }

  var blockTypeConstructors = map[BlockType]func(string) Block {
    Header        : func(line string) Block { return NewHeaderBlock(line) },
    ThematicBreak : func(line string) Block { return NewThematicBreakBlock(line) },
    Code          : func(line string) Block { return NewCodeBlock(line) },
  }

  var processingOrder = []BlockType{Code, Header, ThematicBreak}
  
  for _, blockType := range processingOrder {
    re, _ := blockTypeRegexs[blockType]
    if re.Match([]byte(line)) {
      constructor, _ := blockTypeConstructors[blockType]
      return constructor(line)
    }
  }
  
  return NewParagraphBlock(line)
}

func NewHeaderBlock(line string) *HeaderBlock {
	hashRe := regexp.MustCompile(`#{1,6}`)
  endingRe := regexp.MustCompile(` [# ]+$`)
  headingLevel := len(hashRe.Find([]byte(line)))

  trimmedString := endingRe.ReplaceAllString(strings.TrimLeft(line, "# "), "")
	
  return &HeaderBlock{
		content : trimmedString,
		Level   : headingLevel,
	}
}

func NewCodeBlock(line string) *CodeBlock {
	return &CodeBlock{
		content: line,
    isOpen:  true,
	}
}

func NewParagraphBlock(content string) *ParagraphBlock {
	return &ParagraphBlock{
		content: content,
	}
}

func NewThematicBreakBlock(line string) *ThematicBreakBlock {
  print("tb")
  return &ThematicBreakBlock{}
}


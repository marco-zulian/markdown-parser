package blocks

import (
  "fmt"
  "regexp"
)

type ParagraphBlock struct {
	content string
  isOpen  bool
}

func (paragraph *ParagraphBlock) GetBlockType() BlockType {
	return Paragraph
}

func (paragraph *ParagraphBlock) GetContent() *string {
	return &paragraph.content
}

func (paragraph *ParagraphBlock) String() string {
  return fmt.Sprintf("Paragraph: %s", paragraph.content)
}

func (paragraph *ParagraphBlock) CanConsume(line string) bool {
  if !paragraph.isOpen { return false }

  var blankLineRegex = regexp.MustCompile("^ *$")
  if blankLineRegex.Match([]byte(line)) {
    paragraph.isOpen = false
    return false
  }

  return true
}

func (paragraph *ParagraphBlock) Consume(line string) {
  paragraph.content += "\n" + line 
}

func (paragraph *ParagraphBlock) CanExtend() bool {
  return paragraph.isOpen 
}

func (paragraph *ParagraphBlock) IsOpen() bool {
  return paragraph.isOpen
}

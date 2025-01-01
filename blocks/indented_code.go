package blocks

import (
  "regexp"
  "fmt"
)

type CodeBlock struct {
	content string
  isOpen  bool
}

func (codeBlock CodeBlock) GetBlockType() BlockType {
	return Code
}

func (codeBlock CodeBlock) GetContent() *string {
	return &codeBlock.content
}

func (codeBlock CodeBlock) String() string {
  return fmt.Sprintf("Code: %s", codeBlock.content)
}

func (codeBlock *CodeBlock) CanConsume(line string) bool {
  if !codeBlock.isOpen { return false }

  var canConsumeCodeBlockRe = regexp.MustCompile(`(^ {4,}|\t)|(^[ \t]*$)`)
  if canConsumeCodeBlockRe.Match([]byte(line)) { return true }
 
  codeBlock.Close()
  return false 
}

func (codeBlock *CodeBlock) Consume(line string) {
  emptyLineUpToFourCharactersRe := regexp.MustCompile("^ {0,4}$")
  if emptyLineUpToFourCharactersRe.Match([]byte(line)) {
    codeBlock.content += "\n"
    return
  }

  codeBlock.content += "\n" + line[4:] 
}

func (codeBlock *CodeBlock) CanExtend() bool {
  return codeBlock.isOpen 
}

func (codeBlock *CodeBlock) IsOpen() bool {
  return codeBlock.isOpen 
}

func (codeBlock *CodeBlock) Close() {
  emptyLinesAtEndRe := regexp.MustCompile(`(\n\s*)+$`)
  codeBlock.content = emptyLinesAtEndRe.ReplaceAllString(codeBlock.content, "")
  codeBlock.isOpen = false
}

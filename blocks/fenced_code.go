package blocks

import (
  "fmt"
  "regexp"
)

type FencedCodeBlock struct {
  content         string
  openingFenceLen int
  delimiter       string
  info            string
  isOpen          bool
  tabs            int
}

func (fencedCode *FencedCodeBlock) CanConsume(line string) bool {
  return fencedCode.isOpen
}

func (fencedCode *FencedCodeBlock) Consume(line string) {
  compileString := fmt.Sprintf("^\\s{0,3}%s{%d,}\\s*$", fencedCode.delimiter, fencedCode.openingFenceLen)
  closeFenceRe := regexp.MustCompile(compileString)

  if closeFenceRe.Match([]byte(line)) {
    fencedCode.Close()
    return
  }

  trimRe := regexp.MustCompile(fmt.Sprintf("^\\s{0,%d}", fencedCode.tabs))
  trimmedLine := trimRe.ReplaceAllString(line, "")

  if fencedCode.content == "" {
    fencedCode.content += trimmedLine
  } else {
    fencedCode.content += "\n" + trimmedLine 
  }
}

func (fencedCode *FencedCodeBlock) CanExtend() bool {
  return fencedCode.isOpen
}

func (fencedCode *FencedCodeBlock) String() string {
  return fmt.Sprintf("Fenced code: %s", fencedCode.content)
}

func (fencedCode *FencedCodeBlock) IsOpen() bool {
  return fencedCode.isOpen
}

func (fencedCode *FencedCodeBlock) Close() {
  fencedCode.isOpen = false
}

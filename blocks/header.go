package blocks

import "fmt"

type HeaderBlock struct {
	content string
	Level   int
}

func (header *HeaderBlock) String() string {
  return fmt.Sprintf("Header %d: %s", header.Level, header.content)
}

func (header *HeaderBlock) CanConsume(line string) bool {
  return false
}

func (header *HeaderBlock) Consume(line string) {
  panic("Header consuming line")
}

func (header *HeaderBlock) CanExtend() bool {
  return false
}

func (header *HeaderBlock) IsOpen() bool {
  return false
}

func (header *HeaderBlock) Close() {
  return
}

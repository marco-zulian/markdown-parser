package blocks

import "fmt"

type ThematicBreakBlock struct{}

func (thematicBreak *ThematicBreakBlock) GetBlockType() BlockType {
	return ThematicBreak
}

func (thematicBreak *ThematicBreakBlock) GetContent() *string {
	return nil
}

func (thematicBreak *ThematicBreakBlock) String() string {
  return fmt.Sprintf("Thematic break") 
}

func (thematicBreak *ThematicBreakBlock) CanConsume(line string) bool {
  return false
}

func (thematicBreak *ThematicBreakBlock) Consume(line string) {
  panic("Thematic break consuming line")
}
 
func (thematicBreak *ThematicBreakBlock) CanExtend() bool {
  return false 
}

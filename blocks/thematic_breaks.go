package blocks

type ThematicBreakBlock struct{}

func NewThematicBreakBlock() ThematicBreakBlock {
	return ThematicBreakBlock{}
}

func (thematicBreak ThematicBreakBlock) GetBlockType() BlockType {
	return ThematicBreak
}

func (thematicBreak ThematicBreakBlock) GetContent() *string {
	return nil
}

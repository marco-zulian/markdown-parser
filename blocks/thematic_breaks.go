package blocks

type ThematicBreakBlock struct {
  tokenType TokenType
}

func NewThematicBreakBlock() ThematicBreakBlock {
  return ThematicBreakBlock{
    tokenType: ThematicBreak,
  }
}

func (thematicBreak ThematicBreakBlock) GetTokenType() TokenType {
  return thematicBreak.tokenType
}

func (thematicBreak ThematicBreakBlock) GetContent() *string {
  return nil
}

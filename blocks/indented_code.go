package blocks

type IndentedCodeBlock struct {
  tokenType TokenType
  content   string
}

func NewIntendedCodeBlock(content string) IndentedCodeBlock {
  return IndentedCodeBlock{
    tokenType: IndentedCode,
    content: content,
  }
}

func (indentedCode IndentedCodeBlock) GetTokenType() TokenType {
  return indentedCode.tokenType
}

func (indentedCode IndentedCodeBlock) GetContent() *string {
  return &indentedCode.content
}

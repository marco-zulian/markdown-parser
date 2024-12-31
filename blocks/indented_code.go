package blocks

type IndentedCodeBlock struct {
	content string
}

func NewIntendedCodeBlock(content string) IndentedCodeBlock {
	return IndentedCodeBlock{
		content: content,
	}
}

func (indentedCode IndentedCodeBlock) GetBlockType() BlockType {
	return IndentedCode
}

func (indentedCode IndentedCodeBlock) GetContent() *string {
	return &indentedCode.content
}

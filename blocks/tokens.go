package blocks

type TokenType string

const (
	Paragraph     TokenType = "paragraph"
	Header        TokenType = "header"
  ThematicBreak TokenType = "break"
  IndentedCode  TokenType = "indented_code"
)

type Token interface {
	GetTokenType() TokenType
	GetContent() *string
}

package blocks

type TokenType string

const (
	Paragraph TokenType = "paragraph"
	Header    TokenType = "header"
)

type Token interface {
	GetTokenType() TokenType
	GetContent() *string
}

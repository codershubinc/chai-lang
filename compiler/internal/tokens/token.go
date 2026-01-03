package tokens

// TokenType represents the type of a token.
type TokenType string

const (
	// TOKEN_EOF represents the end of the file.
	TOKEN_EOF TokenType = "EOF"
	// TOKEN_CHAI_SAY represents the 'chai_say' keyword.
	TOKEN_CHAI_SAY TokenType = "CHAI_SAY"
	// TOKEN_STRING represents a string literal.
	TOKEN_STRING TokenType = "STRING"
	// TOKEN_ILLEGAL represents an illegal or unknown token.
	TOKEN_ILLEGAL TokenType = "ILLEGAL"
)

// Token represents a lexical token in the source code.
// It holds the type of the token and its literal value.
type Token struct {
	Type    TokenType
	Literal string
}

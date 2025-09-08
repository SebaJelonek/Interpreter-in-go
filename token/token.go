package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	IF       = "IF"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// delimiters
	SEMICOLON = ";"
	COMMA     = ","

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

var Keywords = map[string]TokenType{
	"let":      LET,
	"function": FUNCTION,
	"if":       IF,
}

package lexer

import (
	"testing"

	token "../token"
)

func TestLexert(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, testCase := range tests {
		lexerToken := lexer.NextToken()

		if testCase.expectedType != lexerToken.Type {
			t.Fatalf("test[%d] - token type wrong\n expected=%q, got=%q", i, testCase.expectedType, lexerToken.Type)
		}

		if testCase.expectedLiteral != lexerToken.Literal {
			t.Fatalf("test[%d] - literal wrong\n expected=%q, got=%q", i, testCase.expectedLiteral, lexerToken.Literal)
		}
	}

}

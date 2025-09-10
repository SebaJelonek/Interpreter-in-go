package lexer

import (
	"testing"

	"github.com/SebaJelonek/Interpreter-in-go/token"
)

func TestLexer(t *testing.T) {
	input := `=+(){},;`
	test := []struct {
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

	for i, testCase := range test {
		lexerToken := lexer.NextToken()

		if testCase.expectedType != lexerToken.Type {
			t.Fatalf("test[%d] - token type wrong\n expected=%q, got=%q, literl=%q", i, testCase.expectedType, lexerToken.Type, lexerToken.Literal)
		}

		if testCase.expectedLiteral != lexerToken.Literal {
			t.Fatalf("test[%d] - literal wrong\n expected=%q, got=%q", i, testCase.expectedLiteral, lexerToken.Literal)
		}
		t.Log("TEST PASSED", i+1)
	}

}
func TestNextToken(t *testing.T) {
	input := `	let five = 5;
				let ten = 10;
				let add = fn(x, y) {
					x + y;
				};
				let result = add(five, ten);
				`
	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, testCase := range test {
		lexerToken := lexer.NextToken()

		if testCase.expectedType != lexerToken.Type {
			t.Fatalf("test[%d] - token type wrong\n expected=%q, got=%q, literl=%q", i, testCase.expectedType, lexerToken.Type, lexerToken.Literal)
		}

		if testCase.expectedLiteral != lexerToken.Literal {
			t.Fatalf("test[%d] - literal wrong\n expected=%q, got=%q", i, testCase.expectedLiteral, lexerToken.Literal)
		}
		t.Log("TEST PASSED", i+1)
	}
}

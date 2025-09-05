package lexer

import (
	"log"

	"github.com/SebaJelonek/Interpreter-in-go/token"
)

type Lexer struct {
	input        string
	position     int  // current char position (index)
	readPosition int  // next char position (index+1)
	ch           byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	log.Println("hello lexer")
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.ch = l.readChar()

	switch rune(l.ch) {
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	}

	return tok
} //this function builds tokens with chars which are comming from readChar()

func (l *Lexer) readChar() byte {
	var char byte
	if l.readPosition >= len(l.input) {
		char = 0
		return char
	}
	char = l.input[l.position]
	log.Println(rune(char))
	//but token takes in string as literal because it would only work on operators
	l.position++
	l.readPosition++

	return char
	//basic char function
	//takes char from current position
	//and returns it
	// additionally moves position and read position by 1

}

func (l *Lexer) peek() byte {

	nextChar := l.input[l.readPosition]
	return nextChar
} //this function peeks into l.readPosition

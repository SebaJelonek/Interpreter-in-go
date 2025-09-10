package lexer

import (
	"log"
	"strings"

	"github.com/SebaJelonek/Interpreter-in-go/token"
)

type Lexer struct {
	input        string
	position     int  // current char position (index) for read
	readPosition int  // next char position (index+1) for peek
	ch           byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	log.Println("hello lexer")
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	log.Println("this is byte at the begining", l.ch)
	l.skipWhiteSpace()
	// this does not work... i will try to fix it today... tomorrow today
	var tok token.Token
	var literal string
	log.Println("this is char:", string(l.ch))
	log.Println("this is byte:", l.ch)

	switch l.ch {
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
		l.readChar()
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
		l.readChar()
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
		l.readChar()
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
		l.readChar()
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
		l.readChar()
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		l.readChar()
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
		l.readChar()
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
		l.readChar()
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default: //the char is one of above chars letter/underscore/number?
		if isLetter(l.ch) || isNumber(l.ch) {
			literal = l.readIdentifier()
			log.Println("the litteral", literal)
			tok = literalToToken(literal)
			log.Println("the token", tok)
			log.Println(l.position)
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
			return tok
		}
	}
	log.Println("the token", tok)
	return tok
} //this function builds tokens with chars which are comming from readChar()

func (l *Lexer) readChar() byte {
	var char byte

	if l.readPosition >= len(l.input) {
		l.ch = 0
		char = 0
		return char
	}

	char = l.input[l.position]
	//but token takes in string as literal because it would only work on operators
	l.position++
	l.readPosition++
	l.ch = char // reading char at the very beginning

	return char
	//basic char function
	//takes char from current position
	//and returns it
	//additionally moves position and read position by 1
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_')
}
func isNumber(char byte) bool {
	return (char >= '0' && char <= '9')
}

func (l *Lexer) readIdentifier() string {
	var literal []string
	// literal = append(literal, string(char)) //<- append the char consumed by nexttoken to slice
	// //so this is not needed
	for isLetter(l.ch) || isNumber(l.ch) /*<-check the first char which comes from */ {
		literal = append(literal, string(l.ch)) // <- we append "old" char, the one read last iteration
		l.ch = l.readChar()                     //<- if this is number or letter we carry on and go back to top
		//if it is not we break out of loop through initial condition...
		// literal = append(literal, string(char)) <- this is an issue... we append the "new" char
	}
	return strings.Join(literal, "") //and return literal
}

func literalToToken(literal string) token.Token {
	var tok token.Token

	val, ok := token.Keywords[literal]
	if ok {
		tok.Literal = literal
		tok.Type = val
	} else if isNumber(literal[0]) {
		tok.Literal = literal
		tok.Type = token.INT
	} else {
		tok.Literal = literal
		tok.Type = token.IDENT
	}

	// now we else
	return tok
}

func (l *Lexer) peek() byte {

	nextChar := l.input[l.readPosition]
	return nextChar
} //this function peeks into l.readPosition

func (l *Lexer) skipWhiteSpace() {
	for l.ch == 32 || l.ch == 9 || l.ch == 10 {
		l.readChar()
	}
}

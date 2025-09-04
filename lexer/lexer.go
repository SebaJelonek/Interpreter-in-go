package lexer

import (
	"log"

	"github.com/SebaJelonek/Interpreter-in-go/token"
)

type Lexer struct {
	Input        string
	Position     int  // current char position (index)
	ReadPosition int  // next char position (index+1)
	ch           byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{Input: input}
	log.Println("hello lexer")
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{Literal: char}
	return tok
} //this function builds tokens with chars which are comming from readChar()

func (l *Lexer) readChar() byte {
	char := l.Input[l.Position]

	log.Println(char)
	//but token takes in string as literal because it would only work on operators
	l.ch = char
	l.Position++
	l.ReadPosition++

	return char
	//basic char function
	//takes char from current position
	//and returns it
	// additionally moves position and read position by 1

}
func (l *Lexer) peek() byte {

	nextChar := l.Input[l.ReadPosition]
	return nextChar
} //this function peeks into l.readposition

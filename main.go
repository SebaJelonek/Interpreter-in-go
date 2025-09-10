package main

import (
	"log"

	"github.com/SebaJelonek/Interpreter-in-go/lexer"
)

func main() {
	log.Println("main")
	lex := mainLexer()
	lex.NextToken()

}
func mainLexer() *lexer.Lexer {
	input := `	let five = 5;
				let ten = 10;
				let add = fn(x, y) {
					x + y;
				};
				let result = add(five, ten);
				`
	l := lexer.New(input)
	return l
}

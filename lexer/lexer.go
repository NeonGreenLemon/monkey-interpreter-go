package lexer

import "github.com/NeonGreenLemon/monkey-interpreter-go/token"

type Lexer struct {
	input string // the actual input
	position int // current position
	readPosition int // the position after the current position
	ch byte // byte representation of the current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // init positon
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	token, err := token.GetToken(l.ch)
	if err != nil {
	  panic(err)
	}
	l.readChar()
	return token
}

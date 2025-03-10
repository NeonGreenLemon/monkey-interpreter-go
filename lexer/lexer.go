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
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LEFT_PARENT, l.ch)
	case ')':
		tok = newToken(token.RIGHT_PARENT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LEFT_BRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHT_BRACE, l.ch)
	case '[':
		tok = newToken(token.LEFT_BRACKET, l.ch)
	case ']':
		tok = newToken(token.RIGHT_BRACKET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

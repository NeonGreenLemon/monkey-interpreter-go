package lexer

import (
	"github.com/NeonGreenLemon/monkey-interpreter-go/token"
)

type Lexer struct {
	input        string // the actual input
	position     int    // current position
	readPosition int    // the position after the current position
	ch           byte   // byte representation of the current character
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	if isLetter(l.ch) {
		tok = token.GetTokenFromString(l.readString())
	} else if isDigit(l.ch) {
		tok = token.GetTokenFromDigit(l.readDigit())
	} else if isTwoChar(l) {
		ch := l.ch
		l.readChar()
		tok = token.GetTwoCharToken(ch, l.ch)
		l.readChar()
	} else {
		tok = token.GetToken(l.ch)
		// go to next character
		l.readChar()
	}

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isTwoChar(l *Lexer) bool {
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			return true
		}
	case '!':
		if l.peekChar() == '=' {
			return true
		}
	}
	return false
}

func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

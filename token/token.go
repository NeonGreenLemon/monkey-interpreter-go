package token

import "errors"

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

func GetToken(ch byte) (Token, error) {
  var tok Token

  switch ch {
  case '=':
    tok = newToken(ASSIGN, ch)
  case ';':
    tok = newToken(SEMICOLON, ch)
  case '(':
    tok = newToken(LEFT_PARENT, ch)
  case ')':
    tok = newToken(RIGHT_PARENT, ch)
  case ',':
    tok = newToken(COMMA, ch)
  case '+':
    tok = newToken(PLUS, ch)
  case '{':
    tok = newToken(LEFT_BRACE, ch)
  case '}':
    tok = newToken(RIGHT_BRACE, ch)
  case '[':
    tok = newToken(LEFT_BRACKET, ch)
  case ']':
    tok = newToken(RIGHT_BRACKET, ch)
  case 0:
    tok.Literal = ""
    tok.Type = EOF
  default:
    return Token{Type: EOF, Literal: ""}, errors.New("could not map char")
  }

  return tok, nil
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

const(
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifier
    IDENTIFIER = "IDENTIFIER"
    INT = "INT"

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    LEFT_PARENT = "("
    RIGHT_PARENT = ")"
    LEFT_BRACKET = "["
    RIGHT_BRACKET = "]"
    LEFT_BRACE = "{"
    RIGHT_BRACE = "}"

    COMMA = ","
    SEMICOLON = ";"

    // Keywords
    LET = "LET"
    FUNCTION = "FUNCTION"
)

package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

constant {
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
}

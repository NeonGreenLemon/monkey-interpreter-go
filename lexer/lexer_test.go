package lexer

import (
	"testing"

	"github.com/NeonGreenLemon/monkey-interpreter-go/token"
)

func TestNextToken(t *testing.T) {
	input := `+={}()[],;`

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
		{token.LEFT_PARENT, "("},
		{token.RIGHT_PARENT, ")"},
		{token.LEFT_BRACKET, "["},
		{token.RIGHT_BRACKET, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got%q",
			i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got%q",
			i, tt.expectedLiteral, tok.Literal)
		}
	}
}

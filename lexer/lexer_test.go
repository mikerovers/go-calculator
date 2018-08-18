package lexer

import (
	"testing"
	"go-calculator/token"
	)

type Test []struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextTokenNumbers(t *testing.T) {
	input := "123\n"

	test := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.INT, "1"},
		{token.INT, "2"},
		{token.INT, "3"},
	}

	l := New(input)

	for i, tt := range test {
		tok := l.nextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenEmpty(t *testing.T) {
	input := "       "

	l := New(input)
	tok := l.nextToken()

	if tok.Type != token.EOF {
		t.Fatalf("TestNextTokenEmpty - tokentype wrong. expected=%q, got=%q", token.EOF, tok.Type)
	}

	if tok.Literal != "" {
		t.Fatalf("TestNextTokenEmpty - literal wrong. expected=%q, got=%q", "``", tok.Literal)
	}
}

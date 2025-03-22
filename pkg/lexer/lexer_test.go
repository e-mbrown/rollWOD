package lexer

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/lexer/samples"
	"github.com/e-mbrown/rollWOD/pkg/token"
)

func TestNextToken(t *testing.T) {
	fl, err := samples.ServeSample("test_tokens")
	if err != nil {
		t.Fatal("Sample `test_tokens.wql` has been altered, deleted or moved. Delete or update test")
	}
	defer fl.Close()

	l, err := NewLexer(fl)
	if err != nil {
		t.Fatalf("Problems creating lexer, most like issue with file. Err = %s", err.Error())
	}

	tests := []struct {
		expType    token.TokenType
		expLiteral []byte
	}{
		{token.ASSIGN, []byte("=")},
		{token.PLUS, []byte("+")},
		{token.LPAREN, []byte("(")},
		{token.RPAREN, []byte(")")},
		{token.LBRACE, []byte("{")},
		{token.RBRACE, []byte("}")},
		{token.SEMICOLON, []byte(";")},
		{token.COMMA, []byte(",")},
		{token.ASTERISK, []byte("*")},
		{token.SEMICOLON, []byte(";")},
		{token.EOF, []byte{0}},
	}

	for i, tt := range tests {

		tok := l.NextToken()

		if tok.Type != tt.expType {
			t.Fatalf("test[%d] - tokentype incorrect. expected=%q got=%q",
				i, tt.expType, tok.Type)
		}

		if string(tok.Literal) != string(tt.expLiteral) {
			t.Fatalf("test[%d] - literal incorrect. expected=%q got=%q",
				i, tt.expLiteral, tok.Literal)
		}
	}

}

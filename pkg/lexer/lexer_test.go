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
		{token.DECLARE, []byte("dcl")},
		{token.IDENT, []byte("input")},
		{token.ASSIGN, []byte("=")},
		{token.INT, []byte("10")},
		{token.SEMICOLON, []byte(";")},
		{token.DECLARE, []byte("dcl")},
		{token.IDENT, []byte("double")},
		{token.ASSIGN, []byte("=")},
		{token.FUNCTION, []byte("fn")},
		{token.LPAREN, []byte("(")},
		{token.IDENT, []byte("x")},
		{token.RPAREN, []byte(")")},
		{token.LBRACE, []byte("{")},
		{token.IDENT, []byte("x")},
		{token.ASTERISK, []byte("*")},
		{token.INT, []byte("2")},
		{token.RBRACE, []byte("}")},
		{token.SEMICOLON, []byte(";")},
		{token.FUNCTION, []byte("fn")},
		{token.IDENT, []byte("customMulti")},
		{token.LPAREN, []byte("(")},
		{token.IDENT, []byte("x")},
		{token.COMMA, []byte(",")},
		{token.IDENT, []byte("y")},
		{token.RPAREN, []byte(")")},
		{token.LBRACE, []byte("{")},
		{token.IDENT, []byte("x")},
		{token.ASTERISK, []byte("*")},
		{token.IDENT, []byte("y")},
		{token.RBRACE, []byte("}")},
		{token.SEMICOLON, []byte(";")},
		{token.DECLARE, []byte("dcl")},
		{token.IDENT, []byte("doub")},
		{token.ASSIGN, []byte("=")},
		{token.IDENT, []byte("double")},
		{token.LPAREN, []byte("(")},
		{token.IDENT, []byte("input")},
		{token.RPAREN, []byte(")")},
		{token.SEMICOLON, []byte(";")},
		{token.DECLARE, []byte("dcl")},
		{token.IDENT, []byte("res")},
		{token.ASSIGN, []byte("=")},
		{token.IDENT, []byte("customMulti")},
		{token.LPAREN, []byte("(")},
		{token.IDENT, []byte("input")},
		{token.COMMA, []byte(",")},
		{token.IDENT, []byte("doub")},
		{token.RPAREN, []byte(")")},
		{token.SEMICOLON, []byte(";")},
		{token.DECLARE, []byte("dcl")},
		{token.IDENT, []byte("input2")},
		{token.ASSIGN, []byte("=")},
		{token.FLOAT, []byte("10.23")},
		{token.SEMICOLON, []byte(";")},
		{token.FLOAT, []byte("11.4")},
		{token.RCARET, []byte(">")},
		{token.INT, []byte("6")},
		{token.MINUS, []byte("-")},
		{token.INT, []byte("7")},
		{token.LCARET, []byte("<")},
		{token.INT, []byte("0")},
		{token.SEMICOLON, []byte(";")},
		{token.INT, []byte("10")},
		{token.N_EQ, []byte("!=")},
		{token.INT, []byte("12")},
		{token.SEMICOLON, []byte(";")},
		{token.IF, []byte("if")},
		{token.TRUE, []byte("true")},
		{token.LBRACE, []byte("{")},
		{token.RETURN, []byte("return")},
		{token.RBRACE, []byte("}")},
		{token.ELSE, []byte("else")},
		{token.LBRACE, []byte("{")},
		{token.RETURN, []byte("return")},
		{token.FALSE, []byte("false")},
		{token.RBRACE, []byte("}")},
		{token.SEMICOLON, []byte(";")},
		{token.EOF, []byte{0}},
	}

	for i, tt := range tests {

		tok := l.NextToken()

		if tok.Type != tt.expType {
			t.Fatalf("test[%d]: %q - tokentype incorrect. expected=%q got=%q",
				i, tok.Literal, tt.expType, tok.Type)
		}

		if string(tok.Literal) != string(tt.expLiteral) {
			t.Fatalf("test[%d] - literal incorrect. expected=%q got=%q",
				i, tt.expLiteral, tok.Literal)
		}
	}

}

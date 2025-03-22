package lexer

import (
	"bufio"
	"os"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type Lexer struct {
	filename string
	reader   *bufio.Reader // might not be needed
	input    []byte
	pos      int
	readPos  int
	ch       byte
}

func NewLexer(input *os.File) (*Lexer, error) {
	var err error
	l := &Lexer{
		filename: input.Name(),
		reader:   bufio.NewReader(input),
	}

	// TODO: This may need to be a scanner, not a reader. The purpose
	// is to read the file & store the bytes. Dont need a delimiter.
	l.input, err = l.reader.ReadBytes(byte(0))
	if err != nil && err.Error() != "EOF"  {
		return nil, err
	}

	l.readChar()
	return l, nil
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case byte('='):
		tok = newToken(token.ASSIGN, l.ch)
	case byte('+'):
		tok = newToken(token.PLUS, l.ch)
	case byte('-'):
		tok = newToken(token.MINUS, l.ch)
	case byte('*'):
		tok = newToken(token.ASTERISK, l.ch)
	case byte(','):
		tok = newToken(token.COMMA, l.ch)
	case byte(';'):
		tok = newToken(token.SEMICOLON, l.ch)
	case byte('('):
		tok = newToken(token.LPAREN, l.ch)
	case byte(')'):
		tok = newToken(token.RPAREN, l.ch)
	case byte('{'):
		tok = newToken(token.LBRACE, l.ch)
	case byte('}'):
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = []byte{0}
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenTyp token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenTyp, Literal: []byte{ch}}
}

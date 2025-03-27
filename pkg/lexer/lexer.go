package lexer

import (
	"bufio"
	"os"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type Lexer struct {
	Filename string
	FileLn   int
	reader   *bufio.Reader // might not be needed
	input    []byte
	pos      int
	readPos  int
	ch       byte
}

func NewLexer(input *os.File) (*Lexer, error) {
	var l *Lexer
	var err error
	if input == nil {
		l = &Lexer{FileLn: 1}
	} else {
		l = &Lexer{
			Filename: input.Name(),
			FileLn:   1,
			reader:   bufio.NewReader(input),
		}
		// TODO: This may need to be a scanner, not a reader. The purpose
		// is to read the file & store the bytes. Dont need a delimiter.
		l.input, err = l.reader.ReadBytes(byte('~'))
		if err != nil && err.Error() != "EOF" {
			return nil, err
		}

		l.readChar()
	}

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

	l.skipWhitespace()

	switch l.ch {
	case byte('='):
		if l.peekChar() == byte('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: []byte{ch, l.ch}}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
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
	case byte('!'):
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.N_EQ, Literal: []byte{ch, l.ch}}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case byte('<'):
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.N_EQ, Literal: []byte{ch, l.ch}}
		} else {
			tok = newToken(token.LCARET, l.ch)
		}
	case byte('>'):
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.N_EQ, Literal: []byte{ch, l.ch}}
		} else {
			tok = newToken(token.RCARET, l.ch)
		}
	case 0:
		tok.Literal = []byte{0}
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdent()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		} else if isNumber(l.ch) {
			var isFloat bool

			tok.Literal, isFloat = l.readNum()
			if isFloat {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}

			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenTyp token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenTyp, Literal: []byte{ch}}
}

func (l *Lexer) readIdent() []byte {
	pos := l.pos
	for isLetter(l.ch) || isNumber(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) readNum() ([]byte, bool) {
	var isFloat bool
	pos := l.pos
	for isNumber(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			isFloat = true
		}
		l.readChar()
	}

	return l.input[pos:l.pos], isFloat
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.FileLn++
		}
		l.readChar()
	}
}

func (l *Lexer) TakeInput(input string) {
	l.input = []byte(input)
	l.readChar()
}

// TODO: Add other allowed identifiers when they are thought of.
func isLetter(ch byte) bool {
	return ch >= 97 && ch <= 122 || ch >= 65 && ch <= 90 ||
		ch == '_'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

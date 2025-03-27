package parser

import (
	"fmt"

	"github.com/e-mbrown/rollWOD/pkg/lexer"
	"github.com/e-mbrown/rollWOD/pkg/token"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

type Parser struct {
	l *lexer.Lexer

	currTok token.Token
	peekTok token.Token

	errs []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l:l, errs: []string{}}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) ParseWQL() *wql.WQLRoot {
	root := &wql.WQLRoot{}
	root.Stmts = []wql.Stmt{}

	for p.currTok.Type != token.EOF {
		stmt := p.parseStmt()
		if stmt != nil {
			root.Stmts = append(root.Stmts, stmt)
		}
		p.nextToken()
	}

	return root
}

func (p *Parser) parseStmt() wql.Stmt{
	switch p.currTok.Type {
	case token.DECLARE:
		return p.parseDeclareStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return nil
	}
}

func (p *Parser) parseDeclareStmt() *wql.DeclareStmt {
	stmt := &wql.DeclareStmt{Token: p.currTok}

	// predict a name
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &wql.Identifier{Token: p.currTok, Val: p.currTok.Literal}
	//predict assign
	if !p.expectPeek(token.ASSIGN){
		return nil
	}

	//TODO ...

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStmt() *wql.ReturnStmt {
	stmt := &wql.ReturnStmt{Token: p.currTok}
	
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func(p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currTok.Type == t
}

func(p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekTok.Type == t
}

func(p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func(p *Parser) Errors() []string {
	return p.errs
}

func(p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("%s.ln %d--expected next token to be %s, got %s instead", p.l.Filename, p.l.FileLn ,t, p.peekTok.Type)
	p.errs = append(p.errs, msg)
}
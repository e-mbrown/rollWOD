package parser

import (
	"fmt"
	"strconv"

	"github.com/e-mbrown/rollWOD/pkg/lexer"
	"github.com/e-mbrown/rollWOD/pkg/token"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	LESSGREATEREQUAL
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParse func() wql.Expr
	infixParse  func(wql.Expr) wql.Expr
)

type Parser struct {
	l *lexer.Lexer

	currTok token.Token
	peekTok token.Token

	preParse map[token.TokenType]prefixParse
	inParse  map[token.TokenType]infixParse
	errs     []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errs: []string{}}

	p.preParse = make(map[token.TokenType]prefixParse)
	registerPrefixFn(p)

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

func (p *Parser) parseStmt() wql.Stmt {
	switch p.currTok.Type {
	case token.DECLARE:
		return p.parseDeclareStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExprStmt()
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
	if !p.expectPeek(token.ASSIGN) {
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

func (p *Parser) parseExprStmt() *wql.ExprStmt {
	stmt := &wql.ExprStmt{Token: p.currTok}

	stmt.Expr = p.parseExpr(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpr(precedence int) wql.Expr {
	prefix := p.preParse[p.currTok.Type]
	if prefix == nil {
		p.noPreParseErr(p.currTok.Type)
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdent() wql.Expr {
	return &wql.Identifier{Token: p.currTok, Val: p.currTok.Literal}
}

func (p *Parser) parseIntLiteral() wql.Expr {
	lit := &wql.IntLiteral{Token: p.currTok}

	val, err := strconv.ParseInt(string(p.currTok.Literal), 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as int", p.currTok.Literal)
		p.errs = append(p.errs, msg)
		return nil
	}

	lit.Val = val
	return lit
}

func (p *Parser) parsePreExpr() wql.Expr {
	expr := &wql.PrefixExpr{
		Token: p.currTok,
		Op: string(p.currTok.Literal),
	}

	p.nextToken()
	expr.Right = p.parseExpr(PREFIX)

	return expr
}

/*HELPER FUNCTIONS*/

func (p *Parser) regInfix(t token.TokenType, fn infixParse) {
	p.inParse[t] = fn
}
func (p *Parser) regPrefix(t token.TokenType, fn prefixParse) {
	p.preParse[t] = fn
}

func registerPrefixFn(p *Parser){
	p.regPrefix(token.IDENT, p.parseIdent)
	p.regPrefix(token.INT, p.parseIntLiteral)
	p.regPrefix(token.BANG, p.parsePreExpr)
	p.regPrefix(token.MINUS, p.parsePreExpr)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currTok.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekTok.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errs
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("%s.ln %d--expected next token to be %s, got %s instead", p.l.Filename, p.l.FileLn, t, p.peekTok.Type)
	p.errs = append(p.errs, msg)
}

func (p *Parser) noPreParseErr(t token.TokenType) {
	msg:= fmt.Sprintf("no prefix parse func for %s found", t)
	p.errs = append(p.errs, msg)
}
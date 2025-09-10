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

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.N_EQ:     EQUALS,
	token.LCARET:   LESSGREATER,
	token.RCARET:   LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

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
	p.inParse = make(map[token.TokenType]infixParse)
	registerFn(p)
	registerQfn(p)

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
	case token.QUERY:
		return p.parseQueryStmt()
	case token.CREATE:
		return p.parseCreateStmt()
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

	p.nextToken()
	stmt.Val = p.parseExpr(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStmt() *wql.ReturnStmt {
	stmt := &wql.ReturnStmt{Token: p.currTok}

	p.nextToken()

	if p.curTokenIs(token.SEMICOLON) {
		return stmt
	}

	stmt.ReturnVal = p.parseExpr(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
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
		p.noPreParseErr(p.currTok)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.inParse[p.peekTok.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdent() wql.Expr {
	return &wql.Identifier{Token: p.currTok, Val: p.currTok.Literal}
}

func (p *Parser) parseBool() wql.Expr {
	return &wql.BoolLiteral{Token: p.currTok, Val: p.curTokenIs(token.TRUE)}
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

func (p *Parser) parseIfExpr() wql.Expr {
	expr := &wql.IfExpr{Token: p.currTok}

	p.nextToken()
	expr.Cond = p.parseExpr(LOWEST)

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expr.IfBlock = p.parseBlockStmt()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expr.ElBlock = p.parseBlockStmt()
	}
	return expr
}

func (p *Parser) parsePreExpr() wql.Expr {
	expr := &wql.PrefixExpr{
		Token: p.currTok,
		Op:    string(p.currTok.Literal),
	}

	p.nextToken()
	expr.Right = p.parseExpr(PREFIX)

	return expr
}

func (p *Parser) parseGroupedExpr() wql.Expr {
	p.nextToken()

	expr := p.parseExpr(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return expr
}

func (p *Parser) parseInExpr(left wql.Expr) wql.Expr {
	expr := &wql.InfixExpr{
		Token: p.currTok,
		Op:    string(p.currTok.Literal),
		Left:  left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expr.Right = p.parseExpr(precedence)

	return expr
}

func (p *Parser) parseFuncLit() wql.Expr {
	fnLit := &wql.FuncLit{Token: p.currTok}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	fnLit.Params = p.parseFuncParams()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	fnLit.Body = p.parseBlockStmt()

	return fnLit
}

func (p *Parser) parseCallExpr(function wql.Expr) wql.Expr {
	expr := &wql.CallExpr{Token: p.currTok, Func: function}
	expr.Args = p.parseCallArgs()
	return expr
}

func (p *Parser) parseCallArgs() []wql.Expr {
	args := []wql.Expr{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpr(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpr(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseBlockStmt() *wql.BlockStmt {
	block := &wql.BlockStmt{Token: p.currTok}
	block.Stmts = []wql.Stmt{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStmt()
		if stmt != nil {
			block.Stmts = append(block.Stmts, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseFuncParams() []*wql.Identifier {
	idents := []*wql.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return idents
	}

	p.nextToken()

	ident := &wql.Identifier{Token: p.currTok, Val: p.currTok.Literal}
	idents = append(idents, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &wql.Identifier{Token: p.currTok, Val: p.currTok.Literal}
		idents = append(idents, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return idents
}

func (p *Parser) parseListLiteral() wql.Expr {
	expr := &wql.ListLiteral{Token: p.currTok}
	expr.Val = p.parseListArgs()
	return expr
}

func (p *Parser) parseListArgs() []wql.Expr {
	args := []wql.Expr{}

	if p.peekTokenIs(token.RBRACKET) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpr(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpr(LOWEST))
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return args
}

/*HELPER FUNCTIONS*/

func (p *Parser) regInfix(t token.TokenType, fn infixParse) {
	p.inParse[t] = fn
}
func (p *Parser) regPrefix(t token.TokenType, fn prefixParse) {
	p.preParse[t] = fn
}

func registerFn(p *Parser) {
	p.regPrefix(token.IDENT, p.parseIdent)
	p.regPrefix(token.INT, p.parseIntLiteral)
	p.regPrefix(token.BANG, p.parsePreExpr)
	p.regPrefix(token.MINUS, p.parsePreExpr)
	p.regPrefix(token.TRUE, p.parseBool)
	p.regPrefix(token.FALSE, p.parseBool)
	p.regPrefix(token.LPAREN, p.parseGroupedExpr)
	p.regPrefix(token.IF, p.parseIfExpr)
	p.regPrefix(token.FUNCTION, p.parseFuncLit)
	p.regPrefix(token.FUNCTION, p.parseFuncLit)
	p.regPrefix(token.LBRACKET, p.parseListLiteral)

	p.regInfix(token.MINUS, p.parseInExpr)
	p.regInfix(token.PLUS, p.parseInExpr)
	p.regInfix(token.SLASH, p.parseInExpr)
	p.regInfix(token.ASTERISK, p.parseInExpr)
	p.regInfix(token.EQ, p.parseInExpr)
	p.regInfix(token.N_EQ, p.parseInExpr)
	p.regInfix(token.LCARET, p.parseInExpr)
	p.regInfix(token.RCARET, p.parseInExpr)
	p.regInfix(token.LPAREN, p.parseCallExpr)

}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currTok.Type == t
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.currTok.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekTok.Type == t
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekTok.Type]; ok {
		return p
	}

	return LOWEST
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

func (p *Parser) noPreParseErr(t token.Token) {
	msg := fmt.Sprintf("no prefix parse func for %s found. Literal=%s", t, string(t.Literal))
	p.errs = append(p.errs, msg)
}

package parser

import (
	"github.com/e-mbrown/rollWOD/pkg/token"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func (p *Parser) parseQueryStmt() *wql.QueryStmt {
	query := &wql.QueryStmt{Token: p.currTok}

	if !p.expectPeek(token.LBRACE){
		return nil
	}

	query.Body = p.parseBlockStmt()

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return query
}

func (p *Parser) parseCreateAlterStmt() *wql.CreateAlterStmt {
	stmt := &wql.CreateAlterStmt{Token: p.currTok}

	p.nextToken()
	
	stmt.Stmt = p.parseExpr(LOWEST)
	
	for p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}



func (p *Parser) parseCreateObjStmt() wql.Expr {
	expr := &wql.ObjStmt{Token: p.currTok}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	expr.Params = p.parseFuncParams()

	if !p.expectPeek(token.VALUES) {
		return nil
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	expr.Args = p.parseCallArgs()

	return expr
}
 

func registerQfn(p *Parser){
	p.regPrefix(token.USER, p.parseCreateObjStmt)
	p.regPrefix(token.CAMPAIGN,p.parseCreateObjStmt)
	p.regPrefix(token.CHARACTER,p.parseCreateObjStmt)
}
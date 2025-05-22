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

func (p *Parser) parseCreateStmt() *wql.CreateStmt {
	stmt := &wql.CreateStmt{Token: p.currTok}

	if !p.expectPeek(token.USER) && !p.expectPeek(token.CHARACTER) && !p.expectPeek(token.CAMPAIGN) {
		return nil
	}

	stmt.Stmt = p.parseUserStmt()
	
	for p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}



func (p *Parser) parseUserStmt() *wql.UserStmt {
	expr := &wql.UserStmt{Token: p.currTok}

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
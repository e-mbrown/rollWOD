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
	create := &wql.CreateStmt{Token: p.currTok}

	obj := &wql.ObjStmt{}
	
	p.nextToken()
	obj.StmtType =  p.currTok.Literal
	create.Stmt = obj

	//Probably diff parse functions
	// switch string(p.currTok.Literal) {
	// case "user":
	// 	create.Stmt = obj
	// case "campaign":
	// case "character":
	// }
	p.nextToken()
	return create
}
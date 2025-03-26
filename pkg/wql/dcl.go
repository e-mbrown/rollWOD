package wql

import "github.com/e-mbrown/rollWOD/pkg/token"



type DeclareStmt struct {
	Token token.Token
	Name  *Identifier
	Val   Expr
}

func (ds *DeclareStmt) stmtNode()          {}
func (ds *DeclareStmt) TokLiteral() []byte { return ds.Token.Literal }

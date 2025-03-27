package wql

import "github.com/e-mbrown/rollWOD/pkg/token"


type ExprStmt struct {
	Token token.Token
	Expr Expr
}

func (ex *ExprStmt) stmtNode() {}
func (ex *ExprStmt) TokenLiteral() []byte {return ex.Token.Literal}
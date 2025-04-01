package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type ExprStmt struct {
	Token token.Token
	Expr  Expr
}

type PrefixExpr struct {
	Token token.Token
	Op    string
	Right Expr
}

func (ex *ExprStmt) String() string       { return Expr.String(ex.Expr) }
func (ex *ExprStmt) stmtNode()            {}
func (ex *ExprStmt) TokenLiteral() []byte { return ex.Token.Literal }

func (pe *PrefixExpr) exprNode()            {}
func (pe *PrefixExpr) TokenLiteral() []byte { return pe.Token.Literal }
func (pe *PrefixExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Op)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

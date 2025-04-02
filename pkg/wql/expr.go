package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type ExprStmt struct {
	Token token.Token
	Expr  Expr
}

func (ex *ExprStmt) String() string       { return Expr.String(ex.Expr) }
func (ex *ExprStmt) stmtNode()            {}
func (ex *ExprStmt) TokenLiteral() []byte { return ex.Token.Literal }


type PrefixExpr struct {
	Token token.Token
	Op    string
	Right Expr
}

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


type InfixExpr struct {
	Token token.Token
	Op    string
	Left Expr
	Right Expr
}

func (ie *InfixExpr) exprNode()            {}
func (ie *InfixExpr) TokenLiteral() []byte { return ie.Token.Literal }
func (ie *InfixExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Op + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

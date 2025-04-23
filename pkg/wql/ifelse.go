package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type IfExpr struct {
	Token token.Token
	Cond Expr
	IfBlock *BlockStmt
	ElBlock	* BlockStmt
}


func (ie *IfExpr) exprNode() {}
func (ie *IfExpr) TokenLiteral()[]byte {return ie.Token.Literal}
func (ie *IfExpr) String() string {
	var output bytes.Buffer

	output.WriteString("if")
	output.WriteString(ie.Cond.String())
	output.WriteString(" ")
	output.WriteString(ie.IfBlock.String())

	if ie.ElBlock != nil {
		output.WriteString(ie.ElBlock.String())
	}
	output.WriteString(";")

	return output.String()
}
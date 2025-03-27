package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)



type DeclareStmt struct {
	Token token.Token
	Name  *Identifier
	Val   Expr
}


func (ds *DeclareStmt) String() string{
	var output bytes.Buffer

	output.WriteString(string(ds.TokenLiteral())+ " ")
	output.WriteString(ds.Name.String())
	output.WriteString(" = ")

	if ds.Val != nil {
		output.WriteString(ds.Val.String())
	}
	output.WriteString(";")

	return output.String()
}

func (ds *DeclareStmt) stmtNode()          {}
func (ds *DeclareStmt) TokenLiteral() []byte { return ds.Token.Literal }

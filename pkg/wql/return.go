package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type ReturnStmt struct {
	Token     token.Token
	ReturnVal Expr
}

func (rs *ReturnStmt) String() string {
	var output bytes.Buffer

	output.WriteString(string(rs.TokenLiteral()) + " ")

	if rs.ReturnVal != nil {
		output.WriteString(rs.ReturnVal.String())
	}

	output.WriteString(";")

	return output.String()
}
func (rs *ReturnStmt) stmtNode()            {}
func (rs *ReturnStmt) TokenLiteral() []byte { return rs.Token.Literal }

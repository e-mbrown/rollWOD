package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type QueryStmt struct {
	Token token.Token
	Body *BlockStmt
}

func(qs *QueryStmt) stmtNode(){}
func(qs *QueryStmt) TokenLiteral() []byte {return qs.Token.Literal}
func(qs *QueryStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(qs.TokenLiteral()))
	out.WriteString(qs.Body.String())

	return out.String()
}
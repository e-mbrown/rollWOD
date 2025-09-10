package wql

import (
	"bytes"
	"strings"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type CreateStmt struct {
	Token token.Token
	Stmt  Expr
}

func (cs *CreateStmt) stmtNode()            {}
func (cs *CreateStmt) TokenLiteral() []byte { return cs.Token.Literal }
func (cs *CreateStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(cs.TokenLiteral()))
	out.WriteString(cs.Stmt.String())

	return out.String()
}

type CreateObjStmt struct {
	Token  token.Token
	Params []*Identifier
	Args   []Expr
}

func (cos *CreateObjStmt) exprNode()            {}
func (cos *CreateObjStmt) TokenLiteral() []byte { return cos.Token.Literal }
func (cos *CreateObjStmt) String() string {
	var out bytes.Buffer

	params := []string{}
	args := []string{}
	for i := 0; i < len(cos.Params); i++ {
		params = append(params, cos.Params[i].String())
		args = append(args, cos.Args[i].String())
	}

	out.WriteString(string(cos.Token.Literal))
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") Values")
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(") ")

	return out.String()
}
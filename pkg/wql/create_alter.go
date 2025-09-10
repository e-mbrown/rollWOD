package wql

import (
	"bytes"
	"strings"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type CreateAlterStmt struct {
	Token token.Token
	Stmt  Expr
}

func (cs *CreateAlterStmt) stmtNode()            {}
func (cs *CreateAlterStmt) TokenLiteral() []byte { return cs.Token.Literal }
func (cs *CreateAlterStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(cs.TokenLiteral()))
	out.WriteString(cs.Stmt.String())

	return out.String()
}

type ObjStmt struct {
	Token  token.Token
	Params []*Identifier
	Args   []Expr
}

func (cos *ObjStmt) exprNode()            {}
func (cos *ObjStmt) TokenLiteral() []byte { return cos.Token.Literal }
func (cos *ObjStmt) String() string {
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
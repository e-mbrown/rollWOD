package wql

import (
	"bytes"

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

func (os *ObjStmt) exprNode()            {}
func (os *ObjStmt) TokenLiteral() []byte { return os.Token.Literal }
func (os *ObjStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(os.Token.Literal))
	out.WriteString("(")
	for i, val := range os.Params {
		out.WriteString(val.String())
		if i != len(os.Params)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(") Values")
	out.WriteString("(")
	for i, val := range os.Args {
		out.WriteString(val.String())
		if i != len(os.Args)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(") ")

	return out.String()
}

package wql

import (
	"bytes"
	"strings"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type CreateStmt struct {
	Token token.Token
	Stmt  ObjExpr
}

func (cs *CreateStmt) stmtNode()            {}
func (cs *CreateStmt) TokenLiteral() []byte { return cs.Token.Literal }
func (cs *CreateStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(cs.TokenLiteral()))
	out.WriteString(cs.Stmt.String())

	return out.String()
}

type Obj interface {
	TokenLiteral() []byte
	String() string
}
type ObjExpr interface {
	Obj
	ObjNode()
}

type UserStmt struct {
	Token  token.Token
	Params [] *Identifier
	Args   [] Expr
}

func (ue UserStmt) ObjNode()              {}
func (ue *UserStmt) ExprNode()            {}
func (ue *UserStmt) TokenLiteral() []byte { return ue.Token.Literal }
func (ue *UserStmt) String() string {
	var out bytes.Buffer

	params := []string{}
	args := []string{}
	for i := 0; i < len(ue.Params); i++ {
		params = append(params, ue.Params[i].String())
		args = append(args, ue.Args[i].String())
	}

	out.WriteString(string(ue.Token.Literal))
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") Values")
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(") ")

	return out.String()
}

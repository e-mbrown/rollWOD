package wql

import (
	"bytes"
	"fmt"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type SelectStmt struct {
	Token token.Token
	Args []Expr
	From FromClause
	Where WhereClause
}

func (ss *SelectStmt) stmtNode()            {}
func (ss *SelectStmt) TokenLiteral() []byte { return ss.Token.Literal }
func (ss *SelectStmt) String() string {
	var out bytes.Buffer

	out.WriteString(string(ss.Token.Literal))
	out.WriteString(" ")
	out.WriteString(ss.From.String())
	out.WriteString(" ")
	out.WriteString(ss.Where.String())

	return out.String()
}

type WhereClause struct {
	Token token.Token
	Clause DeclareStmt
}

func (fc *WhereClause) stmtNode()            {}
func (fc *WhereClause) TokenLiteral() []byte { return fc.Token.Literal }
func (fc *WhereClause) String() string {return fmt.Sprint(fc.Token.Type, " ", fc.Clause.String())}

// Expr or stmt??
type FromClause struct {
	Token token.Token
	Obj Expr
}

func (fc *FromClause) stmtNode()            {}
func (fc *FromClause) TokenLiteral() []byte { return fc.Token.Literal }
func (fc *FromClause) String() string {return fmt.Sprint(fc.Token.Type, " ", fc.Obj.String())}

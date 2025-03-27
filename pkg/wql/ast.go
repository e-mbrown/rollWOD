package wql

import (
	"bytes"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type Node interface {
	TokenLiteral() []byte
	String() string
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type WQLRoot struct {
	Stmts []Stmt
}

func (w *WQLRoot) TokenLiteral() []byte {
	if len(w.Stmts) > 0 {
		return w.Stmts[0].TokenLiteral()
	} else {
		return []byte{}
	}
}

func (w *WQLRoot) String() string {
	var output bytes.Buffer

	for _, s := range w.Stmts {
		output.WriteString(s.String())
	}

	return output.String()
}

type Identifier struct {
	Token token.Token
	Val   []byte
}

func (i *Identifier) exprNode()            {}
func (i *Identifier) String() string       { return string(i.Val) }
func (i *Identifier) TokenLiteral() []byte { return i.Token.Literal }

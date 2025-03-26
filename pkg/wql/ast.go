package wql

import "github.com/e-mbrown/rollWOD/pkg/token"


type Node interface {
	TokLiteral() []byte
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

func (w *WQLRoot) TokLiteral() []byte {
	if len(w.Stmts) > 0 {
		return w.Stmts[0].TokLiteral()
	} else {
		return []byte{}
	}
}


type Identifier struct {
	Token token.Token
	Val   []byte
}

func (i *Identifier) exprNode()          {}
func (i *Identifier) TokLiteral() []byte { return i.Token.Literal }
package wql

import "github.com/e-mbrown/rollWOD/pkg/token"

type IntLiteral struct {
	Token token.Token
	Val   int64
}

func (il *IntLiteral) exprNode()            {}
func (il *IntLiteral) TokenLiteral() []byte { return il.Token.Literal }
func (il *IntLiteral) String() string       { return string(il.Token.Literal) }

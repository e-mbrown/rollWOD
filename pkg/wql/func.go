package wql

import (
	"bytes"
	"strings"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type FuncLit struct {
	Token token.Token
	Params [] *Identifier
	Body *BlockStmt
}

func(fl *FuncLit) exprNode(){}
func(fl *FuncLit) TokenLiteral() []byte {return fl.Token.Literal}
func(fl *FuncLit) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Params {
		params = append(params, p.String())
	}

	out.WriteString(string(fl.TokenLiteral()))
	out.WriteString("(")
	out.WriteString(strings.Join(params,", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}
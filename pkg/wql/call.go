package wql

import (
	"bytes"
	"strings"

	"github.com/e-mbrown/rollWOD/pkg/token"
)

type CallExpr struct {
	Token token.Token
	Func Expr
	Args []Expr
}

func (ce *CallExpr) exprNode() {}
func (ce *CallExpr) TokenLiteral() []byte {return ce.Token.Literal}
func (ce *CallExpr) String() string {
	var out bytes.Buffer

	args := []string{}
		for _, a := range ce.Args {
			args = append(args, a.String())
		}

		out.WriteString(ce.Func.String())
		out.WriteString("(")
		out.WriteString(strings.Join(args, ", "))
		out.WriteString(")")

		return out.String()
}

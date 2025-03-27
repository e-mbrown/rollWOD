package wql_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/token"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func TestString(t *testing.T) {
	root := &wql.WQLRoot{
		Stmts: []wql.Stmt{
			&wql.DeclareStmt{
				Token: token.Token{Type: token.DECLARE, Literal: []byte("dcl")},
				Name: &wql.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: []byte("aVar")},
					Val:   []byte("aVar"),
				},
				Val: &wql.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: []byte("nextVar")},
					Val:   []byte("nextVar"),
				},
			},
		},
	}

	if root.String() != "dcl aVar = nextVar;" {
		t.Errorf("root.String() wrong. got=%q", root.String())
	}
}

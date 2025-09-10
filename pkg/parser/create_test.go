package parser_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/parser"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func getQueryStmt(t *testing.T, psr *parser.Parser) *wql.QueryStmt {
	wqlRoot := psr.ParseWQL()
	parser.CheckParserErrors(t, psr)
	if wqlRoot == nil {
		t.Fatalf("ParseWQL returned nil")
	}

	if len(wqlRoot.Stmts) != 1 {
		t.Fatalf("Expected 1 parsed statements, got: %d", len(wqlRoot.Stmts))
	}

	qs, ok := wqlRoot.Stmts[0].(*wql.QueryStmt)
	if !ok {
		t.Fatalf("wqlRoot.stmt is not an wql.QueryStmt. got=%T", qs)
	}

	return qs
}

func TestCreateStmt(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "create_stmt")
	qs := getQueryStmt(t, psr)

	tests := []struct {
		expectType     string
		expectedParams []string
		expectedArgs   []any
	}{
		{"character", []string{"test"}, []any{"yippie"}},
		{"campaign", []string{"name", "creator", "members"}, []any{"test", 5,}},
		{"campaign", []string{}, []any{}},
		{"user", []string{"user", "ic", "pass"}, []any{"steve", "xdfv", "supastar"}},
		{"user", []string{"ic", "user"}, []any{"awed", "meev"}},
	}

	for i, tt := range tests {
		create, ok := qs.Body.Stmts[i].(*wql.CreateStmt)
		if !ok {
			t.Fatalf("qs.Body.Stmts[%d] is not an wql.ExprStmt. got=%T", i, create)
		}

		objStmt, ok := create.Stmt.(*wql.CreateObjStmt)
		if !ok {
			t.Errorf("expr isn't wql.CreateObjStmt. got=%T", objStmt)
		}

		for i, ident := range tt.expectedParams {
			parser.TestLiteralExpr(t, objStmt.Params[i], ident)
		}
		
		for i, args := range tt.expectedArgs {
			parser.TestLiteralExpr(t, objStmt.Args[i], args)
		}
	}
}

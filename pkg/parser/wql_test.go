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

func TestCreateAlterStmt(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "create_stmt")
	qs := getQueryStmt(t, psr)

	test := []struct {
		token          string
		expectType     string
		expectedParams []string
		expectedArgs   []any
	}{
		{"CREATE", "character", []string{"test"}, []any{"yippie"}},
		{"CREATE", "campaign", []string{"name", "creator", "members"}, []any{"test", 5, []int{1, 2, 3}}},
		{"CREATE", "campaign", []string{"members", "creator", "name"}, []any{[]int{1}, 5, "testing"}},
		{"CREATE", "user", []string{"user", "ic", "pass"}, []any{"steve", "xdfv", "supastar"}},
		{"CREATE", "user", []string{"ic", "user"}, []any{"awed", "meev"}},
		{"ALTER", "character", []string{"test"}, []any{"yippie"}},
		{"ALTER", "campaign", []string{"name", "creator", "members"}, []any{"test", 5, []int{1, 2, 3}}},
		{"ALTER", "user", []string{"user", "ic", "pass"}, []any{"steve", "xdfv", "supastar"}},
		{"ALTER", "user", []string{"ic", "user"}, []any{"awed", "meev"}},
	}

	for i, tt := range test {
		createAlt, ok := qs.Body.Stmts[i].(*wql.CreateAlterStmt)
		if !ok {
			t.Fatalf("qs.Body.Stmts[%d] is not an wql.CreateAlterStmt. got=%T", i, createAlt)
		}

		if string(createAlt.Token.Type) != tt.token {
			t.Fatalf("createAlt.Token is not an %s. got=%s", tt.token, createAlt.Token.Type)
		}

		objStmt, ok := createAlt.Stmt.(*wql.ObjStmt)
		if !ok {
			t.Errorf("expr isn't wql.ObjStmt. got=%T", objStmt)
		}

		if len(objStmt.Params) != len(tt.expectedParams) ||
			len(objStmt.Args) != len(tt.expectedParams) {
			t.Fatalf("Mismatched params/args expected=params:%d, args:%d. got=params:%d, args:%d", len(tt.expectedParams), len(objStmt.Params), len(tt.expectedArgs), len(objStmt.Args))
		}

		for i, ident := range tt.expectedParams {
			parser.TestLiteralExpr(t, objStmt.Params[i], ident)
		}

		for i, args := range tt.expectedArgs {
			parser.TestLiteralExpr(t, objStmt.Args[i], args)
		}
	}
}

func TestSelectStmt(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "select_stmt")
	qs := getQueryStmt(t, psr)

	tests := []struct {
		expectedArgs []string
		fc           string
		expectIdent  string
		expectVal    any
	}{
		{[]string{"test"}, "user", "item", 2},
		{[]string{"test1, test2"}, "user", "list", []int{2, 3}},
		{[]string{"yippie"}, "user", "exclamation", "yahoo"},
	}

	for i, tt := range tests {
		sel, ok := qs.Body.Stmts[i].(*wql.SelectStmt)
		if !ok {
			t.Fatalf("qs.Body.Stmts[%d] is not an wql.CreateAlterStmt. got=%T", i, sel)
		}

		if len(tt.expectedArgs) != len(sel.Args) {
			t.Fatalf("Mismatched Select args not %d. got=%d", len(tt.expectedArgs), len(sel.Args))
		}

		if !parser.TestIdent(t, sel.From.Obj, tt.expectIdent){
			return
		}

		if !parser.TestIdent(t, sel.Where.Clause.Name, tt.expectIdent) {
			return
		}

		if !parser.TestLiteralExpr(t, sel.Where.Clause.Val, tt.expectVal) {
			return
		}
	}
}

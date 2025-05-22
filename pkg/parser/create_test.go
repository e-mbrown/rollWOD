package parser_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/parser"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func TestCreateStmt(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "create_stmt")
	qs := getQueryStmt(t, psr)

	tests := []struct {
		expectType string
	}{
		{"character"},
		{"campaign"},
		{"user"},
	}

	for i, tt := range tests {
		create, ok := qs.Body.Stmts[i].(*wql.CreateStmt)
		if !ok {
			t.Fatalf("qs.Body.Stmts[%d] is not an wql.CreateStmt. got=%T", i, create)
		}

		if string(create.Stmt.TokenLiteral()) != tt.expectType {
			t.Fatalf("create.StmtType is not %s. got=%s", tt.expectType, create.Stmt.String())
		} 
	}
	
}


func getQueryStmt(t *testing.T, psr *parser.Parser) *wql.QueryStmt{
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
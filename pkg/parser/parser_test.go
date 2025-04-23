package parser_test

import (
	"fmt"
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/parser"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func TestDeclare(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "dcl_stmts")

	wqlRoot := psr.ParseWQL()
	parser.CheckParserErrors(t, psr)
	if wqlRoot == nil {
		t.Fatalf("ParseWQL returned nil")
	}
	if len(wqlRoot.Stmts) != 3 {
		t.Fatalf("Expected 3 parsed statements, got: %d", len(wqlRoot.Stmts))
	}

	tests := []struct {
		expectedIdentifier []byte
	}{
		{[]byte("x2")},
		{[]byte("BigGuy")},
		{[]byte("snake_case")},
	}

	for i, tt := range tests {
		stmt := wqlRoot.Stmts[i]
		if !parser.TestDeclareStmt(t, stmt, string(tt.expectedIdentifier)) {
			return
		}
	}
}

func TestReturnStmt(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "return_stmts")
	wqlRoot := psr.ParseWQL()
	parser.CheckParserErrors(t, psr)

	if wqlRoot == nil {
		t.Fatalf("ParseWQL returned nil")
	}
	if len(wqlRoot.Stmts) != 3 {
		t.Fatalf("Expected 3 parsed statements, got: %d", len(wqlRoot.Stmts))
	}

	for _, stmt := range wqlRoot.Stmts {
		returnStmt, ok := stmt.(*wql.ReturnStmt)
		if !ok {
			t.Errorf("stmt not *wql.ReturnStmt. got: %T", stmt)
			continue
		}
		if string(returnStmt.TokenLiteral()) != "return" {
			t.Errorf("returnstmt.TokenLiteral not 'return'. got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentExpr(t *testing.T) {
	p := parser.PrepareStringParseTest(t, "isIdent;")
	wqlRoot := p.ParseWQL()
	parser.CheckParserErrors(t, p)
	if len(wqlRoot.Stmts) != 1 {
		t.Fatalf("root should have 1 statement. got=%d", len(wqlRoot.Stmts))
	}

	stmt, ok := wqlRoot.Stmts[0].(*wql.ExprStmt)
	if !ok {
		t.Fatalf("wqlRoot.Stmt[0] is not a wql.ExprStmt. got:%T", wqlRoot.Stmts[0])
	}

	if !parser.TestIdent(t, stmt.Expr, "isIdent") {
		return
	}
}

func TestIntLit(t *testing.T) {
	p := parser.PrepareStringParseTest(t, "24;")
	wqlRoot := p.ParseWQL()
	parser.CheckParserErrors(t, p)
	if len(wqlRoot.Stmts) != 1 {
		t.Fatalf("root should have 1 statement. got=%d", len(wqlRoot.Stmts))
	}

	stmt, ok := wqlRoot.Stmts[0].(*wql.ExprStmt)
	if !ok {
		t.Fatalf("wqlRoot.Stmt[0] is not a wql.ExprStmt. got:%T", wqlRoot.Stmts[0])
	}

	parser.TestIntLit(t, stmt.Expr, 24)
}
func TestBoolLit(t *testing.T) {
	p := parser.PrepareFileParseTest(t, "bool_stmt")
	wqlRoot := p.ParseWQL()
	parser.CheckParserErrors(t, p)
	if len(wqlRoot.Stmts) != 4 {
		t.Fatalf("root should have 4 statement. got=%d", len(wqlRoot.Stmts))
	}

	tests := []struct {
		expectedIdentifier bool
	}{
		{true},
		{false},
		{true},
		{false},
	}

	for i, tt := range tests {
		var skipExprStmt bool
		stmt := wqlRoot.Stmts[i]
		exprStmt, ok := stmt.(*wql.ExprStmt)
		if !ok {
			dStmt, ok := stmt.(*wql.DeclareStmt)
			if !ok {
				t.Errorf("stmt is not a declare stmt or exprStmt. got=%T", stmt)
				return
			}
			t.Log(dStmt)
			if !parser.TestBoolLit(t, dStmt.Val, tt.expectedIdentifier) {
				return
			}
			skipExprStmt = true
		}

		if !skipExprStmt && !parser.TestBoolLit(t, exprStmt.Expr, tt.expectedIdentifier) {
			return
		}
	}
}

func TestParsingPrefixExpr(t *testing.T) {
	prefixTest := []struct {
		input  string
		op     string
		intVal interface{}
	}{
		{"!45", "!", 45},
		{"-100", "-", 100},
		{"!true", "!", true},
		{"!false", "!", false},
	}

	for _, tt := range prefixTest {
		p := parser.PrepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		parser.CheckParserErrors(t, p)

		if len(root.Stmts) != 1 {
			t.Fatalf("root should have 1 statement. got=%d", len(root.Stmts))
		}

		stmt, ok := root.Stmts[0].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("wql.Stmt[0] is not a wql.ExprStmt. got:%T", root.Stmts[0])
		}

		expr, ok := stmt.Expr.(*wql.PrefixExpr)
		if !ok {
			t.Fatalf("expr not wql.PrefixExpr. got:%T", expr)
		}

		if expr.Op != tt.op {
			t.Errorf("literal.op not %s. got:%s", expr.Op, tt.op)
		}

		if !parser.TestLiteralExpr(t, expr.Right, tt.intVal) {
			return
		}
	}
}

func TestParsingInfixExpr(t *testing.T) {
	infixTest := []struct {
		input string
		lVal  interface{}
		op    string
		rVal  interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 <= 5;", 5, "<=", 5},
		{"true <= true;", true, "<=", true},
		{"true == true;", true, "==", true},
		{"false != true;", false, "!=", true},
	}

	for _, tt := range infixTest {
		p := parser.PrepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		parser.CheckParserErrors(t, p)

		if len(root.Stmts) != 1 {
			t.Fatalf("root should have 1 statement. got=%d", len(root.Stmts))
		}

		stmt, ok := root.Stmts[0].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("wql.Stmt[0] is not a wql.ExprStmt. got:%T", root.Stmts[0])
		}

		if !parser.TestInfixExpr(t, stmt.Expr, tt.lVal, tt.op, tt.rVal) {
			return
		}
	}
}

func TestOpPrecedenceParse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a + b * c + d / e -f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == false",
			"((3 + (4 * 5)) == false)",
		},
		{
			"true == 3 * 1 + 4 * 5",
			"(true == ((3 * 1) + (4 * 5)))",
		},
		{
			"true == 3 * 1 + 4 * 5",
			"(true == ((3 * 1) + (4 * 5)))",
		},
		{
			"2/(5 * 5)",
			"(2 / (5 * 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		p := parser.PrepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		parser.CheckParserErrors(t, p)

		res := root.String()

		if res != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, res)
		}
	}
}

func TestParseIfExpr(t *testing.T) {
	p := parser.PrepareFileParseTest(t, "if_stmt")

	wqlRoot := p.ParseWQL()
	parser.CheckParserErrors(t, p)

	tests := []struct {
		expectCondLeft  string
		op              string
		expectCondRight string
		expectifBlock   []string
		expectelBlock   string
	}{
		{"x", "<", "y", []string{"x"}, ""},
		{"x", "!=", "y", []string{"x"}, "y"},
	}
	fmt.Println(wqlRoot.Stmts)
	if len(wqlRoot.Stmts) != len(tests) {
		t.Fatalf("Wql body does not contain %d stmts. got=%d", len(tests), len(wqlRoot.Stmts))
	}

	for i, tt := range tests {
		stmt, ok := wqlRoot.Stmts[i].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("wqlRoot.stmt[%d] is not an wql.ExprStmt. got=%T", i, stmt)
		}

		expr, ok := stmt.Expr.(*wql.IfExpr)
		if !ok {
			t.Fatalf("expr is not an wql.IfStmt. got=%T", expr)
		}

		fmt.Println(expr.String())
		if !parser.TestInfixExpr(
			t, expr.Cond, tt.expectCondLeft, tt.op, tt.expectCondRight) {
			return
		}

		if len(expr.IfBlock.Stmts) != len(tt.expectifBlock) {
			t.Fatalf("expr Ifblock stmt len is not %d. got=%d", len(tt.expectifBlock), len(expr.IfBlock.Stmts))
		}

		// Expand to check stmts

		if expr.ElBlock != nil {
			if tt.expectelBlock == "" {
				t.Fatalf("expr.Elblock stmt was not nil. got=%+v", expr.ElBlock)
			}

			if len(expr.ElBlock.Stmts) != len(tt.expectelBlock) {
				t.Fatalf("expr Elblock stmt len is not %d. got=%d", len(tt.expectelBlock), len(expr.ElBlock.Stmts))
			}
		}

	}

}

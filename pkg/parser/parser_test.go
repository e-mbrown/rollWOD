package parser_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/parser"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

type testInfix struct {
	left  interface{}
	op    string
	right interface{}
}

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
		expectIdent string
		expectVal   interface{}
	}{
		{"x", 5},
		{"BigGuy", true},
		{"snake_case", "y"},
	}

	for i, tt := range tests {
		stmt := wqlRoot.Stmts[i]

		if !parser.TestDeclareStmt(t, stmt, tt.expectIdent) {
			return
		}

		val := stmt.(*wql.DeclareStmt).Val
		if !parser.TestLiteralExpr(t, val, tt.expectVal) {
			return
		}
	}
}

func TestQueryBlock(t *testing.T) {
	psr := parser.PrepareFileParseTest(t, "query_block")
	wqlRoot := psr.ParseWQL()
	parser.CheckParserErrors(t, psr)
	if wqlRoot == nil {
		t.Fatalf("ParseWQL returned nil")
	}
	if len(wqlRoot.Stmts) != 2 {
		t.Fatalf("Expected 2 parsed statements, got: %d", len(wqlRoot.Stmts))
	}

	tests := []struct {
		body       int
		Expectstmt []string
	}{
		{0, []string{"10"}},
		{2, []string{"x", "5"}},
	}

	for i, tt := range tests {
		query, ok := wqlRoot.Stmts[i].(*wql.QueryStmt)
		if !ok {
			t.Fatalf("wqlRoot.stmt[%d] is not an wql.QueryStmt. got=%T", i, query)
		}

		for i, tstmt := range tt.Expectstmt{
			body, ok := query.Body.Stmts[i].(*wql.ExprStmt)
			if !ok {
				t.Fatalf("body.stmt[%d] is not an wql.ExprStmt. got=%T", i, body)
			}

			if body.String() != tstmt {
				return
			}
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

	tests := []struct {
		expectVal interface{}
	}{
		{nil},
		{5},
		{"Jim"},
	}

	for i, tt := range tests {
		returnStmt, ok := wqlRoot.Stmts[i].(*wql.ReturnStmt)
		if !ok {
			t.Errorf("stmt not *wql.ReturnStmt. got: %T", wqlRoot.Stmts[i])
			continue
		}

		if string(returnStmt.TokenLiteral()) != "return" {
			t.Errorf("returnstmt.TokenLiteral not 'return'. got %q", returnStmt.TokenLiteral())
		}

		if !parser.TestLiteralExpr(t, returnStmt.ReturnVal, tt.expectVal) {
			return
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
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}

	for _, tt := range tests {
		p := parser.PrepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		parser.CheckParserErrors(t, p)

		res := root.String()

		if res != tt.expected {
			t.Errorf("expected=%q, got=%q \n", tt.expected, res)
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

func TestFuncLitParse(t *testing.T) {
	p := parser.PrepareFileParseTest(t, "func_lit")

	wqlRoot := p.ParseWQL()
	parser.CheckParserErrors(t, p)

	tests := []struct {
		params []string
		body   int
		left   string
		op     string
		right  string
	}{
		{[]string{"x", "y"}, 1, "x", "+", "y"},
	}

	if len(wqlRoot.Stmts) != len(tests) {
		t.Fatalf("Wql body does not contain %d stmts. got=%d", len(tests), len(wqlRoot.Stmts))
	}

	for i, tt := range tests {
		stmt, ok := wqlRoot.Stmts[i].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("wqlRoot.stmt[%d] is not an wql.ExprStmt. got=%T", i, stmt)
		}

		fn, ok := stmt.Expr.(*wql.FuncLit)
		if !ok {
			t.Fatalf("expr is not an wql.FuncLit. got=%T", fn)
		}

		if len(fn.Params) != len(tt.params) {
			t.Fatalf("FuncLit does not have %d params. got=%d", len(tt.params), len(fn.Params))
		}

		for i, ident := range tt.params {
			parser.TestLiteralExpr(t, fn.Params[i], ident)
		}

		if len(fn.Body.Stmts) != tt.body {
			t.Fatalf("FuncLit body len is not %d. got=%d", tt.body, len(fn.Body.Stmts))
		}

		body, ok := fn.Body.Stmts[i].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("body.stmt[%d] is not an wql.ExprStmt. got=%T", i, body)
		}

		parser.TestInfixExpr(t, body.Expr, tt.left, tt.op, tt.right)

	}
}

func TestFuncParamParse(t *testing.T) {
	tests := []struct {
		input        string
		expectParams []string
	}{
		{input: "func() {}", expectParams: []string{}},
		{input: "func(x) {}", expectParams: []string{"x"}},
		{input: "func(x, y, z) {}", expectParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		p := parser.PrepareStringParseTest(t, tt.input)
		wqlRoot := p.ParseWQL()
		parser.CheckParserErrors(t, p)

		stmt := wqlRoot.Stmts[0].(*wql.ExprStmt)
		fn := stmt.Expr.(*wql.FuncLit)

		if len(fn.Params) != len(tt.expectParams) {
			t.Fatalf("FuncLit does not have %d params. got=%d", len(tt.expectParams), len(fn.Params))
		}

		for i, ident := range tt.expectParams {
			parser.TestLiteralExpr(t, fn.Params[i], ident)
		}
	}
}

func TestCallExprParse(t *testing.T) {
	p := parser.PrepareFileParseTest(t, "call_stmt")
	wqlRoot := p.ParseWQL()
	parser.CheckParserErrors(t, p)

	tests := []struct {
		ident string
		args  []interface{}
	}{
		{ident: "add", args: []interface{}{1, testInfix{2, "*", 3}, testInfix{4, "+", 5}}},
	}

	if len(wqlRoot.Stmts) != len(tests) {
		t.Fatalf("WqlRoot.stmt does not contain %d stmts. got=%d\n", len(tests), len(wqlRoot.Stmts))
	}

	for i, tt := range tests {
		stmt, ok := wqlRoot.Stmts[i].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("wqlRoot.stmt[%d] is not an wql.ExprStmt. got=%T", i, stmt)
		}

		call, ok := stmt.Expr.(*wql.CallExpr)
		if !ok {
			t.Fatalf("expr is not an wql.CallExpr. got=%T", call)
		}

		if !parser.TestIdent(t, call.Func, tt.ident) {
			return
		}

		if len(call.Args) != len(tt.args) {
			t.Fatalf("FuncLit body len is not %d. got=%d", len(tt.args), len(call.Args))
		}

		for i, arg := range tt.args {
			switch v := arg.(type) {
			case testInfix:
				parser.TestInfixExpr(t, call.Args[i], v.left, v.op, v.right)
			default:
				parser.TestLiteralExpr(t, call.Args[i], arg)
			}
		}
	}
}

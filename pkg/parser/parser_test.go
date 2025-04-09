package parser_test

import (
	"fmt"
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/lexer"
	"github.com/e-mbrown/rollWOD/pkg/parser"
	"github.com/e-mbrown/rollWOD/pkg/samples"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func TestDeclare(t *testing.T) {
	psr := prepareFileParseTest(t, "dcl_stmts")

	wqlRoot := psr.ParseWQL()
	checkParserErrors(t, psr)
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
		if !testDeclareStmt(t, stmt, string(tt.expectedIdentifier)) {
			return
		}
	}
}

func TestReturnStmt(t *testing.T) {
	psr := prepareFileParseTest(t, "return_stmts")
	wqlRoot := psr.ParseWQL()
	checkParserErrors(t, psr)

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
	p := prepareStringParseTest(t, "isIdent;")
	wqlRoot := p.ParseWQL()
	checkParserErrors(t, p)
	if len(wqlRoot.Stmts) != 1 {
		t.Fatalf("root should have 1 statement. got=%d", len(wqlRoot.Stmts))
	}

	stmt, ok := wqlRoot.Stmts[0].(*wql.ExprStmt)
	if !ok {
		t.Fatalf("wqlRoot.Stmt[0] is not a wql.ExprStmt. got:%T", wqlRoot.Stmts[0])
	}

	if !testIdent(t, stmt.Expr, "isIdent"){
		return
	}
}

func TestIntLit(t *testing.T) {
	p := prepareStringParseTest(t, "24;")
	wqlRoot := p.ParseWQL()
	checkParserErrors(t, p)
	if len(wqlRoot.Stmts) != 1 {
		t.Fatalf("root should have 1 statement. got=%d", len(wqlRoot.Stmts))
	}

	stmt, ok := wqlRoot.Stmts[0].(*wql.ExprStmt)
	if !ok {
		t.Fatalf("wqlRoot.Stmt[0] is not a wql.ExprStmt. got:%T", wqlRoot.Stmts[0])
	}

	testIntLit(t, stmt.Expr, 24)
}
func TestBoolLit(t *testing.T) {
	p := prepareFileParseTest(t, "bool_stmt")
	wqlRoot := p.ParseWQL()
	checkParserErrors(t, p)
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
			if !ok{
				t.Errorf("stmt is not a declare stmt or exprStmt. got=%T", stmt)
				return
			}
			t.Log(dStmt)
			if !testBoolLit(t, dStmt.Val, tt.expectedIdentifier) {
				return
			}
			skipExprStmt = true
		}

		if !skipExprStmt && !testBoolLit(t, exprStmt.Expr, tt.expectedIdentifier) {
			return
		}
	}
}

func TestParsingPrefixExpr(t *testing.T) {
	prefixTest := []struct {
		input  string
		op     string
		intVal int64
	}{
		{"!45", "!", 45},
		{"-100", "-", 100},
	}

	for _, tt := range prefixTest {
		p := prepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		checkParserErrors(t, p)

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
		if !testIntLit(t, expr.Right, tt.intVal) {
			return
		}
	}
}

func TestParsingInfixExpr(t *testing.T) {
	infixTest := []struct {
		input string
		lVal int64
		op string
		rVal int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 <= 5;", 5, "<=", 5},
	}

	for _, tt := range infixTest {
		p := prepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		checkParserErrors(t, p)

		if len(root.Stmts) != 1 {
			t.Fatalf("root should have 1 statement. got=%d", len(root.Stmts))
		}

		stmt, ok := root.Stmts[0].(*wql.ExprStmt)
		if !ok {
			t.Fatalf("wql.Stmt[0] is not a wql.ExprStmt. got:%T", root.Stmts[0])
		}

		if !testInfixExpr(t, stmt.Expr, tt.lVal, tt.op, tt.rVal){
			return
		}
	}
}

func TestOpPrecedenceParse(t *testing.T) {
	tests := []struct {
		input string
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
	}

	for _, tt := range tests {
		p := prepareStringParseTest(t, tt.input)
		root := p.ParseWQL()
		checkParserErrors(t, p)

		res := root.String()

		if res != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, res)
		}
	}
}

/* HELPERS */

func prepareFileParseTest(t *testing.T, fn string) *parser.Parser {
	fl, err := samples.ServeSample(fn)
	if err != nil {
		t.Fatal("Sample `test_tokens.wql` has been altered, deleted or moved. Delete or update test")
	}
	defer fl.Close()

	l, err := lexer.NewLexer(fl)
	if err != nil {
		t.Fatalf("Problems creating lexer, most like issue with file. Err = %s", err.Error())
	}

	psr := parser.NewParser(l)
	return psr
}

func prepareStringParseTest(t *testing.T, input string) *parser.Parser {
	l, err := lexer.NewLexer(nil)
	if err != nil {
		t.Fatalf("Problems creating lexer, most like issue with file. Err = %s", err.Error())
	}

	l.TakeInput(input)
	return parser.NewParser(l)
}

func testLiteralExpr(
	t *testing.T,
	expr wql.Expr,
	expect interface{},
) bool {
	switch v := expect.(type) {
	case int:
		return testIntLit(t, expr, int64(v))
	case int64:
		return testIntLit(t, expr, v)
	case string:
		return testIdent(t, expr, v)
	}
	t.Errorf("type of expr not handled. got=%T", expr)
	return false
}

func testInfixExpr(
	t *testing.T,
	expr wql.Expr,
	left interface{},
	op string,
	right interface{},
) bool {
	opExpr, ok := expr.(*wql.InfixExpr)
	if !ok{
		t.Errorf("expr isn't wql.OperatorExpr. got=%T(%s)", expr, expr)
	}

	if !testLiteralExpr(t, opExpr.Left, left) {
		return false
	}

	if opExpr.Op != op {
		t.Errorf("expr.Op isn't '%s'. got=%q", op, opExpr.Op)
	}

	if !testLiteralExpr(t, opExpr.Right, right) {
		return false
	}

	return true
}

func testDeclareStmt(t *testing.T, s wql.Stmt, name string) bool {
	if string(s.TokenLiteral()) != "dcl" {
		t.Errorf("Token literal not 'dcl'. got=%q", s.TokenLiteral())
		return false
	}

	dclStmt, ok := s.(*wql.DeclareStmt)
	if !ok {
		t.Errorf("current stmt isn't dcl stmt. got: %T", s)
	}

	return testLiteralExpr(t, dclStmt.Name, name)
}

func testIntLit(t *testing.T, il wql.Expr, val int64) bool {
	intLit, ok := il.(*wql.IntLiteral)
	if !ok {
		t.Errorf("il not wql.IntLiteral, got=%d", il)
		return false
	}

	if intLit.Val != val {
		t.Errorf("intLit.Val not %d. got=%d", val, intLit.Val)
		return false
	}

	str := fmt.Sprintf("%d", val)
	if string(intLit.TokenLiteral()) != str {
		t.Errorf("intLit.Val not %s. got=%s", str, string(intLit.TokenLiteral()))
		return false
	}

	return true
}

func testBoolLit(t *testing.T, il wql.Expr, val bool) bool {
	bLit, ok := il.(*wql.BoolLiteral)
	if !ok {
		t.Errorf("il not wql.BoolLiteral, got=%T", il)
		return false
	}

	if bLit.Val != val {
		t.Errorf("intLit.Val not %t. got=%t", val, bLit.Val)
		return false
	}

	str := fmt.Sprintf("%t", val)
	if string(bLit.TokenLiteral()) != str {
		t.Errorf("intLit.Val not %s. got=%s", str, string(bLit.TokenLiteral()))
		return false
	}

	return true
}

func testIdent(t *testing.T, expr wql.Expr, val string) bool {
	ident, ok := expr.(*wql.Identifier)
	if !ok {
		t.Errorf("expr isn't *wql.Identifier, got=%T", expr)
		return false
	}

	if string(ident.Val) != val {
		t.Errorf("ident.Val != %s. got=%s", val, string(ident.Val))
		return false
	}

	if string(ident.TokenLiteral()) != val {
		t.Errorf("ident.Val != %s. got=%s", val, string(ident.TokenLiteral()))
		return false
	}


	return true
}

func checkParserErrors(t *testing.T, psr *parser.Parser) {
	errors := psr.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}


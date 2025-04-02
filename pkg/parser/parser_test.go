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

	ident, ok := stmt.Expr.(*wql.Identifier)
	if !ok {
		t.Fatalf("wqlRoot.Stmt[0] is not a wql.ExprStmt. got:%T", wqlRoot.Stmts[0])
	}

	if string(ident.Val) != "isIdent" {
		t.Errorf("ident.Val is not:%s. got:%s", "isIdent", string(ident.Val))
	}

	if string(ident.TokenLiteral()) != "isIdent" {
		t.Errorf("ident.TokenLiteral is not:%s. got:%s", "isIdent", string(ident.TokenLiteral()))
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

	literal, ok := stmt.Expr.(*wql.IntLiteral)
	if !ok {
		t.Fatalf("expr not wql.IntLiteral. got:%T", literal)
	}

	if literal.Val != 24 {
		t.Errorf("literal.Val not %d. got:%d", 24, literal.Val)
	}
	if string(literal.TokenLiteral()) != "24" {
		t.Errorf("literal.TokenLiteral not %s, got:%s", "24", literal.TokenLiteral())
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

		expr, ok := stmt.Expr.(*wql.InfixExpr)
		if !ok {
			t.Fatalf("expr not wql.InfixExpr. got:%T", expr)
		}

		if !testIntLit(t, expr.Left, tt.lVal) {
			return
		}

		if expr.Op != tt.op {
			t.Errorf("literal.op not %s. got:%s", expr.Op, tt.op)
		}

		if !testIntLit(t, expr.Right, tt.rVal) {
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

func testDeclareStmt(t *testing.T, s wql.Stmt, name string) bool {
	if string(s.TokenLiteral()) != "dcl" {
		t.Errorf("Token literal not 'dcl'. got=%q", s.TokenLiteral())
		return false
	}

	dclStmt, ok := s.(*wql.DeclareStmt)
	if !ok {
		t.Errorf("current stmt isn't dcl stmt. got: %T", s)
	}

	if string(dclStmt.Name.Val) != name {
		t.Errorf("dclStmt name val not: '%s'. got=%s", name, dclStmt.Name.Val)
	}

	if string(dclStmt.Name.TokenLiteral()) != name {
		t.Errorf("dclStmt name token literal not: '%s'. got=%q, %s", name, dclStmt.Name.TokenLiteral(), string(dclStmt.TokenLiteral()))
	}
	return true
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
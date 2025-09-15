package parser

import (
	"fmt"
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/lexer"
	"github.com/e-mbrown/rollWOD/pkg/samples"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func PrepareFileParseTest(t *testing.T, fn string) *Parser {
	fl, err := samples.ServeSample(fn)
	if err != nil {
		t.Fatalf(
			"Sample `%s` may have been altered, deleted, moved or may not exist.Check test name and file",fn)
	}
	defer fl.Close()

	l, err := lexer.NewLexer(fl)
	if err != nil {
		t.Fatalf("Problems creating lexer, most like issue with file. Err = %s", err.Error())
	}

	psr := NewParser(l)
	return psr
}

func PrepareStringParseTest(t *testing.T, input string) *Parser {
	l, err := lexer.NewLexer(nil)
	if err != nil {
		t.Fatalf("Problems creating lexer, most like issue with file. Err = %s", err.Error())
	}

	l.TakeInput(input)
	return NewParser(l)
}

func TestLiteralExpr(
	t *testing.T,
	expr wql.Expr,
	expect any,
) bool {
	switch v := expect.(type) {
	case int:
		return TestIntLit(t, expr, int64(v))
	case int64:
		return TestIntLit(t, expr, v)
	case string:
		return TestIdent(t, expr, v)
	case bool:
		return TestBoolLit(t, expr, v)
	case []any:
		return TestListItems(t, expr, v)
	case []string:
		return TestListItems(t,expr, v)
	case []int:
		return TestListItems(t,expr, v)
	case nil:
		return expr == v
	}
	t.Errorf("type of expr not handled. got=%T", expr)
	return false
}

func TestInfixExpr(
	t *testing.T,
	expr wql.Expr,
	left interface{},
	op string,
	right interface{},
) bool {
	opExpr, ok := expr.(*wql.InfixExpr)
	if !ok {
		t.Errorf("expr isn't wql.OperatorExpr. got=%T(%s)", expr, expr)
	}

	if !TestLiteralExpr(t, opExpr.Left, left) {
		return false
	}

	if opExpr.Op != op {
		t.Errorf("expr.Op isn't '%s'. got=%q", op, opExpr.Op)
	}

	if !TestLiteralExpr(t, opExpr.Right, right) {
		return false
	}

	return true
}

func TestDeclareStmt(t *testing.T, s wql.Stmt, name string) bool {
	if string(s.TokenLiteral()) != "dcl" {
		t.Errorf("Token literal not 'dcl'. got=%q", s.TokenLiteral())
		return false
	}

	dclStmt, ok := s.(*wql.DeclareStmt)
	if !ok {
		t.Errorf("current stmt isn't dcl stmt. got: %T", s)
	}

	return TestLiteralExpr(t, dclStmt.Name, name)
}

func TestIntLit(t *testing.T, il wql.Expr, val int64) bool {
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

func TestBoolLit(t *testing.T, il wql.Expr, val bool) bool {
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

func TestIdent(t *testing.T, expr wql.Expr, val string) bool {
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

func TestListItems[T any](t *testing.T, expr wql.Expr, val []T) bool {
	list, ok := expr.(*wql.ListLiteral)
		if !ok {
			t.Fatalf("expr is not an wql.ListLiteral. got=%T", list)
		}

		
		for i, arg := range val {
			TestLiteralExpr(t, list.Val[i], arg)
		}

		return true
}

func CheckParserErrors(t *testing.T, psr *Parser) {
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

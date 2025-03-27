package parser_test

import (
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/lexer"
	"github.com/e-mbrown/rollWOD/pkg/parser"
	"github.com/e-mbrown/rollWOD/pkg/samples"
	"github.com/e-mbrown/rollWOD/pkg/wql"
)

func TestDeclare(t *testing.T) {
	psr := prepareParseTest(t, "dcl_stmts")
	
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
		if !testDeclareStmt(t, stmt, string(tt.expectedIdentifier)){
			return
		}
	}
}

func TestReturnStmt(t *testing.T){
	psr := prepareParseTest(t, "return_stmts")
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
		if string(returnStmt.TokenLiteral()) != "return"{
			t.Errorf("returnstmt.TokenLiteral not 'return'. got %q", returnStmt.TokenLiteral())
		}
	}
}


func prepareParseTest(t *testing.T, fn string) *parser.Parser {
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
	if err != nil {
		t.Fatalf("Problems creating lexer, most like issue with file. Err = %s", err.Error())
	}

	return psr
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

	if string(dclStmt.Name.Val) != name{
		t.Errorf("dclStmt name val not: '%s'. got=%s", name, dclStmt.Name.Val)
	}

	if string(dclStmt.Name.TokenLiteral()) != name {
		t.Errorf("dclStmt name token literal not: '%s'. got=%q, %s", name, dclStmt.Name.TokenLiteral(), string(dclStmt.TokenLiteral()))
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
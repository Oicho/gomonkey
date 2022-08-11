package parser

import (
	"gomonkey/ast"
	"gomonkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let xy = 51;", "xy", 51},
		{"let x = 0;", "x", 0},
		// TODO {"let x = -5;", "x", -5},

		{"let y = true;", "y", true},
		{"let f = false;", "f", false},
		//{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program, _ := p.ParseProgram()

		//checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return  5;", 5},
		{"return  51;", 51},
		{"return  0;", 0},
		// TODO {"return  -5;", -5},

		{"return  true;", true},
		{"return  false;", false},
		//{"return  y;", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program, _ := p.ParseProgram()

		//checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not *ast.returnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
		if testLiteralExpression(t, returnStmt.Value, tt.expectedValue) {
			return
		}

	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
		//	case string:
		//		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(
	t *testing.T,
	exp ast.Expression,
	expected int64,
) bool {
	ie, err := exp.(*ast.IntegerExpression)
	if !err {
		t.Errorf("ie not *ast.IntegerExpression. got=%T", ie)
		return false
	}

	if ie.Value != expected {
		t.Errorf("ie Value not %d. got=%d", expected, ie.Value)
		return false
	}
	return true
}

func testBooleanLiteral(
	t *testing.T,
	exp ast.Expression,
	expected bool,
) bool {
	ie, err := exp.(*ast.BooleanExpression)
	if !err {
		t.Errorf("be not *ast.BooleanExpression. got=%T", ie)
		return false
	}

	if ie.Value != expected {
		t.Errorf("be Value not %t. got=%t", expected, ie.Value)
		return false
	}
	return true
}

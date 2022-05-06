package lexer

import (
	"errors"
	"testing"
)

type TestResult struct {
	Result string
}

func TestReadString(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
		expectedError  error
	}{
		{"\"This\"", "This", nil},
		{"\"s\"", "s", nil},
		{"\"This is a string\"", "This is a string", nil},
		{"\"asdat", "", errors.New("Reach the end of input without reaching a \"")},
	}
	for i, test := range tests {
		l := New(test.input)
		s, _ := l.ReadString()
		if s != test.expectedOutput {
			t.Fatalf("tests[%d] - string output wrong, expected=%q, got =%q", i, test.expectedOutput, s)
		}
		//if err == nil test.expectedError != nil ||  {
		//	t.Fatalf("tests[%d] - error output wrong, expected=%q, got =%q", i, test.expectedError, err)
		//}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
		expectedError  error
	}{
		{"13", "13", nil},
		{"1", "1", nil},
		{"143  12", "143", nil},
		{"9912+1", "9912", nil},
		{"019", "019", nil},
	}
	for i, test := range tests {
		l := New(test.input)
		s, _ := l.ReadInt()
		if s != test.expectedOutput {
			t.Fatalf("tests[%d] - string output wrong, expected=%q, got =%q", i, test.expectedOutput, s)
		}
		//if err == nil test.expectedError != nil ||  {
		//	t.Fatalf("tests[%d] - error output wrong, expected=%q, got =%q", i, test.expectedError, err)
		//}
	}
}

package lexer

import (
	"errors"
	"gomonkey/token"
)

var name token.TokenType

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespaces()
	switch l.ch {
	case '+':
		t.Literal = "+"
		t.Type = token.PLUS
	case '-':
		t.Literal = "-"
		t.Type = token.MINUS
	case '/':
		t.Literal = "/"
		t.Type = token.DIVIDE
	case '*':
		t.Literal = "*"
		t.Type = token.ASTERISK
	case '!':
		l.PeekChar()
		if l.ch == '=' {
			l.ReadChar()
			t.Literal = "!="
			t.Type = token.NOT_EQ
		} else {
			t.Literal = "!"
			t.Type = token.BANG
		}
	case '=':
		l.PeekChar()
		if l.ch == '=' {
			l.ReadChar()
			t.Literal = "=="
			t.Type = token.EQ
		} else {
			t.Literal = "="
			t.Type = token.ASSIGN
		}
	case '<':
		t.Literal = "<"
		t.Type = token.LT
	case '>':
		t.Literal = ">"
		t.Type = token.GT
	case ',':
		t.Literal = ","
		t.Type = token.COMMA
	case ';':
		t.Literal = ";"
		t.Type = token.SEMICOLON
	case ':':
		t.Literal = ":"
		t.Type = token.COLON
	case '(':
		t.Literal = "("
		t.Type = token.LPAREN
	case ')':
		t.Literal = ")"
		t.Type = token.RPAREN
	case '{':
		t.Literal = "{"
		t.Type = token.LBRACE
	case '}':
		t.Literal = "}"
		t.Type = token.RBRACE
	case '[':
		t.Literal = "["
		t.Type = token.LBRACKET
	case ']':
		t.Literal = "]"
		t.Type = token.RBRACKET
	case '"':
		s, err := l.ReadString()
		if err != nil {
			panic(err)
		}
		t.Literal = s
		t.Type = token.STRING
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isDigit(l.ch) {
			s, err := l.ReadInt()
			if err != nil {
				panic(err)
			}
			t.Literal = s
			t.Type = token.INT
		} else if isLetter(l.ch) {
			// Either a keywords
			s, err := l.ReadIdentifier()
			if err != nil {
				panic(err)
			}
			t.Type = token.LookupIdent(s)
			t.Literal = s
		} else {
			t.Literal = string(l.ch)
			t.Type = token.ILLEGAL
		}
	}
	l.ReadChar()
	return t
}

func New(input string) *Lexer {
	l := &Lexer{input, 0, 0, 0}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) PeekChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
}

func (l *Lexer) ReadString() (string, error) {
	l.ReadChar()
	startPosition := l.position
	for l.ch != '"' && l.ch != 0 {
		l.ReadChar()
	}
	if l.ch == '"' {
		return l.input[startPosition:l.position], nil
	} else {
		return "", errors.New("reach the end of input without reaching a \"")
	}
}

func (l *Lexer) ReadInt() (string, error) {
	startPosition := l.position
	for {
		l.PeekChar()
		if isDigit(l.ch) {
			l.ReadChar()
		} else {
			break
		}
	}
	return l.input[startPosition:l.readPosition], nil
}

func (l *Lexer) ReadIdentifier() (string, error) {
	startPosition := l.position
	for {
		l.PeekChar()
		if isDigit(l.ch) || isLetter(l.ch) || l.ch == '-' || l.ch == '_' {
			l.ReadChar()
		} else {
			break
		}
	}
	return l.input[startPosition:l.readPosition], nil
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

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
	case '"':
		s, err := l.ReadString()
		if err != nil {
			panic(err)
		}
		t.Literal = s
		t.Type = token.STRING
	}
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
		l.ReadChar()
		return l.input[startPosition:l.position], nil
	} else {
		return "", errors.New("reach the end of input without reaching a \"")
	}
}

func (l *Lexer) ReadInt() (string, error) {
	startPosition := l.position
	for {
		l.PeekChar()
		if l.ch >= '0' && l.ch <= '9' {
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

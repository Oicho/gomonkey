package ast

import (
	"bytes"
	"gomonkey/token"
	"strconv"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Right    Expression
	Operator string
}

func (ls *InfixExpression) expressionNode()      {}
func (ls *InfixExpression) TokenLiteral() string { return ls.Token.Literal }
func (ls *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(strconv.Itoa(int(ls.Value)))

	return out.String()
}

type IntegerExpression struct {
	Token token.Token
	Value int64
}

func (ls *IntegerExpression) expressionNode()      {}
func (ls *IntegerExpression) TokenLiteral() string { return ls.Token.Literal }
func (ls *IntegerExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(strconv.Itoa(int(ls.Value)))

	return out.String()
}

type StringExpression struct {
	Token token.Token
	Value string
}

func (ls *StringExpression) expressionNode()      {}
func (ls *StringExpression) TokenLiteral() string { return ls.Token.Literal }
func (ls *StringExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Value)

	return out.String()
}

type BooleanExpression struct {
	Token token.Token
	Value bool
}

func (ls *BooleanExpression) expressionNode()      {}
func (ls *BooleanExpression) TokenLiteral() string { return ls.Token.Literal }
func (ls *BooleanExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	if ls.Value {
		out.WriteString("true")
	} else {
		out.WriteString("false")
	}

	return out.String()
}

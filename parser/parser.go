package parser

import (
	"errors"
	"gomonkey/ast"
	"gomonkey/lexer"
	"gomonkey/token"
	"strconv"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	// check for last token ?
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) ParseExpression() (ast.Expression, error) {
	/*
		exp := ast.Expression()
		for {
			if p.currentToken.Type == token.SEMICOLON {
				return exp, nil
			} else if p.currentToken.Type == token.EOF {
				return nil, errors.New("Unexpected " + string(p.currentToken.Type))
			} else {
				// check fort str int or operator
			}
		}
	*/
	p.NextToken()
	if p.currentToken.Type == token.INT {
		i, err := strconv.Atoi(p.currentToken.Literal)
		ie := &ast.IntegerExpression{Token: p.currentToken, Value: int64(i)}
		p.NextToken()

		return ie, err
	} else if p.currentToken.Type == token.STRING {
		ie := &ast.StringExpression{Token: p.currentToken, Value: p.currentToken.Literal}
		p.NextToken()
		return ie, nil
	}

	return nil, nil
}

func (p *Parser) ParseLetStatement() (*ast.LetStatement, error) {
	ls := &ast.LetStatement{Token: p.currentToken}
	p.NextToken()
	// check  we have identifier token
	if p.currentToken.Type != token.IDENT {
		// TODO check this
		return nil, errors.New("Unexpected " + string(p.currentToken.Type))
	}
	ls.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.NextToken()
	if p.currentToken.Type != token.ASSIGN {
		// TODO check this
		return nil, errors.New("Unexpected " + string(p.currentToken.Type))
	}
	val, err := p.ParseExpression()
	if val == nil {
		return nil, err
	}
	ls.Value = val
	// TODO this should be checked at expression level
	if p.currentToken.Type != token.SEMICOLON {
		// TODO check this
		return nil, errors.New("Unexpected " + string(p.currentToken.Type))
	}
	return ls, nil
}

func (p *Parser) ParseReturnStatement() (*ast.ReturnStatement, error) {
	rt := &ast.ReturnStatement{Token: p.currentToken}
	val, err := p.ParseExpression()
	if val != nil {
		return nil, err
	}
	rt.Value = val
	// TODO this should be checked at expression level
	if p.currentToken.Type != token.SEMICOLON {
		// TODO check this
		return nil, errors.New("Unexpected " + string(p.currentToken.Type))
	}
	return rt, nil
}

func (p *Parser) ParseStatement() (ast.Statement, error) {
	switch p.currentToken.Type {
	case token.EOF:
		return nil, nil
	case token.LET:
		return p.ParseLetStatement()
	case token.RETURN:
		return p.ParseReturnStatement()
	}
	return nil, errors.New("Unexpected " + string(p.currentToken.Type))
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	statement, _ := p.ParseStatement()
	// TODO
	program.Statements = append(program.Statements, statement)

	return program, nil
}

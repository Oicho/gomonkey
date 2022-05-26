package parser

import (
	"errors"
	"gomonkey/ast"
	"gomonkey/lexer"
	"gomonkey/token"
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
	return nil, nil
}

// TODO add error
func (p *Parser) ParseLetStatement() (ast.LetStatement, error) {
	ls := ast.LetStatement{Token: p.currentToken}
	p.NextToken()
	// check  we have identifier token
	ls.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.NextToken()
	if p.currentToken.Type != token.ASSIGN {
		// TODO check this
		return ls, errors.New("Unexpected " + string(p.currentToken.Type))
	}
	for {
		if p.currentToken.Type == token.SEMICOLON {
			break
		} else if p.currentToken.Type == token.EOF {
			panic("Unexpected EOF")
		} else {
			val, err := p.ParseExpression()
			if val == nil {
				return ls, err
			}
			ls.Value = val
		}
	}
	return ls, nil
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	switch p.currentToken.Literal {
	case token.EOF:
		return program, nil
	case token.LET:
		_, err := p.ParseLetStatement()
		if err != nil {
			return program, err
		}
		//program.Statements = append(program.Statements, (ast.Statement)statement)
	}
	return program, nil
}

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

	currentToken    token.Token
	peekToken       token.Token
	prefixFunctions map[token.TokenType]func() (ast.Expression, error)
	infixFunctions  map[token.TokenType]func(ast.Expression) (ast.Expression, error)
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
	p.prefixFunctions = make(map[token.TokenType]func() (ast.Expression, error))
	p.prefixFunctions[token.IDENT] = p.parseIdentifier
	p.prefixFunctions[token.INT] = p.parseIntegerLiteral
	p.prefixFunctions[token.STRING] = p.parseStringLiteral
	p.prefixFunctions[token.TRUE] = p.parseBoolean
	p.prefixFunctions[token.FALSE] = p.parseBoolean
	/*
		p.prefixFunctions(token.BANG, p.parsePrefixExpression)
		p.prefixFunctions(token.MINUS, p.parsePrefixExpression)
		p.prefixFunctions(token.LPAREN, p.parseGroupedExpression)
		p.prefixFunctions(token.IF, p.parseIfExpression)
		p.prefixFunctions(token.FUNCTION, p.parseFunctionLiteral)
		p.prefixFunctions(token.LBRACKET, p.parseArrayLiteral)
		p.prefixFunctions(token.LBRACE, p.parseHashLiteral)

		p.infixParseFns = make(map[token.TokenType]infixParseFn)
		p.registerInfix(token.PLUS, p.parseInfixExpression)
		p.registerInfix(token.MINUS, p.parseInfixExpression)
		p.registerInfix(token.SLASH, p.parseInfixExpression)
		p.registerInfix(token.ASTERISK, p.parseInfixExpression)
		p.registerInfix(token.EQ, p.parseInfixExpression)
		p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
		p.registerInfix(token.LT, p.parseInfixExpression)
		p.registerInfix(token.GT, p.parseInfixExpression)

		p.registerInfix(token.LPAREN, p.parseCallExpression)
		p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	*/
	return p
}

func (p *Parser) ParseExpression() (ast.Expression, error) {
	/*
		Is operator ? -> prefix
		Is bool int Ident or str
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
	} else if p.currentToken.Type == token.TRUE {
		ie := &ast.BooleanExpression{Token: p.currentToken, Value: true}
		p.NextToken()
		return ie, nil
	} else if p.currentToken.Type == token.FALSE {
		ie := &ast.BooleanExpression{Token: p.currentToken, Value: false}
		p.NextToken()
		return ie, nil
	}

	return nil, nil
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	id := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.NextToken()
	return id, nil
}

func (p *Parser) parseIntegerLiteral() (ast.Expression, error) {
	i, err := strconv.Atoi(p.currentToken.Literal)
	if err != nil {
		return nil, err
	}
	ie := &ast.IntegerExpression{Token: p.currentToken, Value: int64(i)}
	p.NextToken()
	return ie, nil
}

func (p *Parser) parseStringLiteral() (ast.Expression, error) {
	ie := &ast.StringExpression{Token: p.currentToken, Value: p.currentToken.Literal}
	p.NextToken()
	return ie, nil
}

func (p *Parser) parseBoolean() (ast.Expression, error) {
	if p.currentToken.Type == token.TRUE {
		be := &ast.BooleanExpression{Token: p.currentToken, Value: true}
		p.NextToken()
		return be, nil
	} else {
		be := &ast.BooleanExpression{Token: p.currentToken, Value: false}
		p.NextToken()
		return be, nil
	}
}

func (p *Parser) parseInfixExpression(left ast.Expression) (ast.Expression, error) {
	//right :=
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
	if err != nil {
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
	for {
		statement, err := p.ParseStatement()
		if err == nil {
			return nil, err
		}
		if statement == nil {
			break
		}
		program.Statements = append(program.Statements, statement)
	}

	return program, nil
}

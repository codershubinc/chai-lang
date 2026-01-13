package main

import (
	"strconv"
)

// 3. THE PARSER
// The Parser takes tokens and builds an AST (Abstract Syntax Tree).
type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram is the entry point. It parses statements until EOF.
func (p *Parser) ParseProgram() []Node {
	var statements []Node
	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.nextToken()
	}
	return statements
}

func (p *Parser) parseStatement() Node {
	switch p.curToken.Type {
	case PRINT:
		return p.parsePrintStmt()
	case IDENT:
		// If it starts with a name, it MUST be an assignment (in our simple language)
		// e.g. "x = 10"
		if p.peekToken.Type == ASSIGN {
			return p.parseAssignStmt()
		}
		// If it's just "x", that's an expression, but we don't support bare expressions as statements yet.
		return nil
	default:
		return nil
	}
}

func (p *Parser) parsePrintStmt() *PrintStmt {
	// We are at 'print'. Move to the expression.
	p.nextToken()

	stmt := &PrintStmt{Value: p.parseExpression()}

	return stmt
}

func (p *Parser) parseAssignStmt() *AssignStmt {
	// We are at 'x'.
	stmt := &AssignStmt{Name: p.curToken.Literal}

	p.nextToken() // Move to '='
	p.nextToken() // Move to the value (e.g. '10')

	stmt.Value = p.parseExpression()

	return stmt
}

// parseExpression handles values and math: "10", "x", "x + 5"
func (p *Parser) parseExpression() Expr {
	// 1. Parse the Left side (it must be a number or variable)
	var left Expr

	if p.curToken.Type == INT {
		val, _ := strconv.Atoi(p.curToken.Literal)
		left = &NumberExpr{Value: val}
	} else if p.curToken.Type == IDENT {
		left = &VarExpr{Name: p.curToken.Literal}
	} else {
		return nil // Error: expected number or variable
	}

	// 2. Check if there is an operator (like +) following it
	// e.g. we just parsed "x", now we see "+"
	if p.peekToken.Type == PLUS {
		p.nextToken() // Move to '+'
		op := p.curToken.Literal

		p.nextToken()                // Move to the Right side
		right := p.parseExpression() // Recursively parse the right side

		return &BinaryExpr{Left: left, Op: op, Right: right}
	}

	return left
}

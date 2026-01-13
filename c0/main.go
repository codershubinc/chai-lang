package main

import (
	"fmt"
	"strings"
)

// 1. TOKEN DEFINITIONS
// These are the types of words our language understands.
type TokenType string

const (
	INT    TokenType = "INT"    // Numbers like 10, 5
	PLUS   TokenType = "PLUS"   // +
	ASSIGN TokenType = "ASSIGN" // =
	IDENT  TokenType = "IDENT"  // Variable names like x, y, result
	PRINT  TokenType = "PRINT"  // The keyword 'print'
	EOF    TokenType = "EOF"    // End of file
)

// Token represents a single word in the source code
type Token struct {
	Type    TokenType
	Literal string
}

// String makes the token print nicely for debugging
func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %q)", t.Type, t.Literal)
}

func main() {
	// The code we want to test
	code := `
		x = 10
		y = 20
		print x + y 
		print x + 5 + y + 8 +y
		print x + 5
		print "Hello, World!"
	`

	fmt.Println("Code to analyze:")
	fmt.Println(code)

	// Debug: Print tokens
	fmt.Println("\nTokens found:")
	lDebug := NewLexer(code)
	var fullStatement strings.Builder
	for {
		tok := lDebug.NextToken()
		if tok.Type == EOF {
			break
		}
		if !(tok.Type == IDENT) || !(tok.Type == PRINT) {
			fullStatement.WriteString(tok.String())
			continue
		}
		fullStatement.WriteString("\n" + tok.String())

	}
	fmt.Println(fullStatement.String())
	// 1. Lexer
	l := NewLexer(code)

	// 2. Parser
	p := NewParser(l)

	fmt.Println("Parsed program", p)
	program := p.ParseProgram()

	fmt.Println("Program ", program)

	fmt.Println("\nAST (Abstract Syntax Tree):")
	for _, stmt := range program {
		fmt.Println(stmt)
	}

	fmt.Println("\nOutput:")
	// 3. Evaluator
	for _, stmt := range program {
		Evaluate(stmt)
	}
}

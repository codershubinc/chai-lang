package ast

import (
	"myLang/compiler/internal/tokens"
)

// Node is the base interface for everything in the AST (Abstract Syntax Tree).
// All nodes in the AST must implement this interface.
type Node interface {
	TokenLiteral() string
}

// Statement is a specific type of Node that represents a statement in the language.
// Statements do not produce a value (e.g., 'print "hello"').
type Statement interface {
	Node
	statementNode()
}

// Program is the root node of the AST.
// It contains a list of statements that make up the program.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the token associated with the first statement in the program.
// If the program has no statements, it returns an empty string.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// ChaiSayStatement represents a chai_say statement in the AST.
// Example: chai_say "hello"
type ChaiSayStatement struct {
	Token tokens.Token // The 'chai_say' token
	Value string       // The string value to be printed
}

// statementNode is a marker method to implement the Statement interface.
func (cs *ChaiSayStatement) statementNode() {}

// TokenLiteral returns the literal value of the token associated with the chai_say statement.
func (cs *ChaiSayStatement) TokenLiteral() string { return cs.Token.Literal }

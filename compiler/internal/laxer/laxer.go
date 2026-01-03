package laxer

import (
	"myLang/compiler/internal/tokens"
)

// Lexer represents a lexical analyzer that tokenizes the input source code.
type Lexer struct {
	input        string
	position     int  // current char position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// NewLexer creates a new Lexer instance for the given input string.
// It initializes the lexer by reading the first character.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input and advances the position.
// If the end of the input is reached, it sets ch to 0 (ASCII code for NUL).
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken returns the next token from the input.
// It skips whitespace and identifies tokens based on the current character.
func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token

	l.skipWhitespace()

	switch l.ch {
	case '"':
		tok.Type = tokens.TOKEN_STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = tokens.TOKEN_EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			if tok.Literal == "chai_say" {
				tok.Type = tokens.TOKEN_CHAI_SAY
			} else {
				tok.Type = tokens.TOKEN_ILLEGAL // We only know 'chai_say' for now
			}
			return tok // Return early to avoid skipping the next char
		} else {
			tok = tokens.Token{Type: tokens.TOKEN_ILLEGAL, Literal: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

// readString reads a string literal enclosed in double quotes.
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

// readIdentifier reads an identifier (a sequence of letters).
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace skips over any whitespace characters (spaces, tabs, newlines).
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter checks if a byte is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

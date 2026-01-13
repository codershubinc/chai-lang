package main

// 2. THE LEXER
// The Lexer takes source code and turns it into tokens.
type Lexer struct {
	input        string // the source code being lexed
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

// readChar gives us the next character and advances our position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" signifies end of file
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NewLexer creates a new Lexer for the given input string
// and initializes it by reading the first character.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Read the first character to initialize l.ch
	return l
}

// NextToken looks at the current character and decides what token it is.
// It acts as a switchboard:
// 1. Skips any whitespace.
// 2. Checks if the character is a known symbol (like '=' or '+').
// 3. If it's a letter, it reads the whole word to see if it's a keyword or variable.
// 4. If it's a digit, it reads the whole number.
// 5. Returns the resulting Token and advances to the next character.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = Token{Type: ASSIGN, Literal: string(l.ch)}
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.ch)}
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	default:
		if l.isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// Check if it's a keyword like "print"
			if tok.Literal == "print" {
				tok.Type = PRINT
			} else {
				tok.Type = IDENT
			}
			return tok // Return early because readIdentifier already advanced us
		} else if l.isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// Unknown character
			tok = Token{Type: "ILLEGAL", Literal: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

// .Helper functions for the lexer

// skipWhitespace advances the lexer's position past any non-meaningful characters.
// It skips over spaces, tabs, newlines, and carriage returns because our language
// doesn't use whitespace for logic (unlike Python).
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readIdentifier reads a sequence of letters to form an identifier or keyword.
// It saves the current position, advances until it finds a non-letter character,
// and returns the slice of the input string that represents the word (e.g., "x", "print", "myVar").
func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a sequence of digits to form an integer literal.
// It works exactly like readIdentifier but looks for digits '0'-'9'.
// It returns the number as a string (e.g., "10", "42").
func (l *Lexer) readNumber() string {
	position := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks if a character is allowed in an identifier.
// We allow lowercase 'a'-'z', uppercase 'A'-'Z', and the underscore '_'.
// This defines what characters can start or be part of a variable name.
func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if a character is a numeric digit '0' through '9'.
// This is used to identify the start and body of integer literals.
func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

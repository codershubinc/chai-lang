package main

import (
	"fmt"
	"myLang/compiler/internal/evaluator"
	"myLang/compiler/internal/laxer"
	"myLang/compiler/internal/parser"
	"os"
)

func main() {

	// 1. Check for file argument
	// Ensure that a filename is provided as a command-line argument.
	// check file extension is *.chai or not

	if len(os.Args[1]) < 6 || os.Args[1][len(os.Args[1])-4:] != ".chi" {
		fmt.Println("Error: File extension must be .chi")
		return
	}

	filename := os.Args[1]
	// Read the content of the file.
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	code := string(bytes)

	// 2. The Pipeline
	// Create a new lexer to tokenize the input code.
	l := laxer.NewLexer(code)
	// Create a new parser to parse the tokens into an AST.
	p := parser.NewParser(l)
	// Parse the program to get the AST.
	program := p.ParseProgram()

	// 3. Run it
	// Evaluate the AST to execute the program.
	evaluator.Eval(program)
}

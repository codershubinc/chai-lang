package main

import (
	"fmt"
	"log"
	"myLang/compiler/internal/compiler"
	laxer "myLang/compiler/internal/lexer"
	"myLang/compiler/internal/parser"
	"myLang/compiler/internal/vm"
	"os"
)

func main() {
	// Remove timestamps for cleaner logs
	log.SetFlags(0)

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

	log.Println("")
	// Parse the program to get the AST.
	program := p.ParseProgram()

	// 3. Compile to bytecode
	// Create a new compiler instance.
	comp := compiler.New()

	// Compile the AST into bytecode.
	err = comp.Compile(program)
	if err != nil {
		fmt.Printf("Compilation error: %v\n", err)
		return
	}

	// Get the compiled bytecode.
	bc := comp.Bytecode()

	// Save bytecode to file
	bytecodeFile := filename[:len(filename)-4] + ".bc"
	err = bc.WriteToBinaryFile(bytecodeFile)
	if err != nil {
		fmt.Printf("Error writing bytecode file: %v\n", err)
		return
	}
	fmt.Printf("Bytecode written to: %s\n", bytecodeFile)

	// 4. Run the bytecode in the VM
	// Create a new VM with the bytecode.
	machine := vm.New(bc)

	// Execute the bytecode.
	err = machine.Run()
	if err != nil {
		fmt.Printf("Runtime error: %v\n", err)
		return
	}
}

package main

import (
	"fmt"
	"myLang/compiler/internal/bytecode"
	"myLang/compiler/internal/vm"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: runner <bytecode-file>")
		return
	}

	// Check file extension
	filename := os.Args[1]
	if len(filename) < 4 || filename[len(filename)-3:] != ".bc" {
		fmt.Println("Error: File extension must be .bc")
		return
	}

	// Load bytecode from file
	bc, err := bytecode.ReadFromBinaryFile(filename)
	if err != nil {
		fmt.Printf("Error reading bytecode file: %v\n", err)
		return
	}

	// Create and run VM
	machine := vm.New(bc)
	err = machine.Run()
	if err != nil {
		fmt.Printf("Runtime error: %v\n", err)
		return
	}
}

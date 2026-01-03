package evaluator

import (
	"fmt"
	"myLang/compiler/internal/ast"
)

// Eval evaluates the given AST node.
// It traverses the AST and executes the statements.
func Eval(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		// Evaluate each statement in the program
		for _, stmt := range node.Statements {
			Eval(stmt)
		}
	case *ast.ChaiSayStatement:
		// Execute the chai_say statement by printing the value to stdout
		fmt.Println(node.Value)
	}
}

package main

import "fmt"

// 4. THE EVALUATOR
// The Evaluator walks the AST and executes the code.

// variables holds the state of our program (memory)
var variables = make(map[string]int)

func Evaluate(node Node) int {
    switch n := node.(type) {
    
    // Statements
    case *AssignStmt:
        val := Evaluate(n.Value)
        variables[n.Name] = val
        return val
        
    case *PrintStmt:
        val := Evaluate(n.Value)
        fmt.Println(val)
        return val
        
    // Expressions
    case *BinaryExpr:
        left := Evaluate(n.Left)
        right := Evaluate(n.Right)
        if n.Op == "+" {
            return left + right
        }
        return 0
        
    case *NumberExpr:
        return n.Value
        
    case *VarExpr:
        return variables[n.Name]
    }
    return 0
}
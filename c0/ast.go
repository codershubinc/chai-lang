package main

import "fmt"

// Node is the basic interface for everything in our tree
type Node interface {
    String() string
}

// -- STATEMENTS (Commands) --

// AssignStmt: x = 10
type AssignStmt struct {
    Name  string // The variable name (e.g., "x")
    Value Expr   // The value being assigned (e.g., 10)
}
func (s *AssignStmt) String() string {
    return fmt.Sprintf("%s = %s", s.Name, s.Value.String())
}

// PrintStmt: print x
type PrintStmt struct {
    Value Expr
}
func (s *PrintStmt) String() string {
    return fmt.Sprintf("print %s", s.Value.String())
}

// -- EXPRESSIONS (Values) --

// Expr is an interface for anything that produces a value
type Expr interface {
    Node
}

// NumberExpr: 10
type NumberExpr struct {
    Value int
}
func (e *NumberExpr) String() string {
    return fmt.Sprintf("%d", e.Value)
}

// VarExpr: x
type VarExpr struct {
    Name string
}
func (e *VarExpr) String() string {
    return e.Name
}

// BinaryExpr: x + 5
type BinaryExpr struct {
    Left  Expr   // x
    Op    string // +
    Right Expr   // 5
}
func (e *BinaryExpr) String() string {
    return fmt.Sprintf("(%s %s %s)", e.Left.String(), e.Op, e.Right.String())
}
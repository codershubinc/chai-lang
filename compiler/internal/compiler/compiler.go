package compiler

import (
	"encoding/binary"
	"myLang/compiler/internal/ast"
	"myLang/compiler/internal/bytecode"
)

// Compiler compiles AST nodes into bytecode
type Compiler struct {
	instructions bytecode.Instructions
	constants    []string
}

// New creates a new Compiler instance
func New() *Compiler {
	return &Compiler{
		instructions: bytecode.Instructions{},
		constants:    []string{},
	}
}

// Compile compiles the AST into bytecode
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, stmt := range node.Statements {
			err := c.Compile(stmt)
			if err != nil {
				return err
			}
		}
	case *ast.ChaiSayStatement:
		// Add the string constant to the constants pool
		constantIndex := c.addConstant(node.Value)

		// Emit OpConstant instruction with the constant index
		c.emit(bytecode.OpConstant, constantIndex)

		// Emit OpPrint instruction to print the value
		c.emit(bytecode.OpPrint)
	}

	return nil
}

// Bytecode returns the compiled bytecode
func (c *Compiler) Bytecode() *bytecode.Bytecode {
	return &bytecode.Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// addConstant adds a constant to the constants pool and returns its index
func (c *Compiler) addConstant(obj string) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

// emit emits a bytecode instruction with optional operands
func (c *Compiler) emit(op bytecode.Opcode, operands ...int) {
	ins := make([]byte, 1)
	ins[0] = byte(op)

	// Add operands (for now, we only support one 2-byte operand)
	for _, operand := range operands {
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(operand))
		ins = append(ins, buf...)
	}

	c.instructions = append(c.instructions, ins...)
}

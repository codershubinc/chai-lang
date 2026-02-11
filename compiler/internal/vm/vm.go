package vm

import (
	"encoding/binary"
	"fmt"
	"myLang/compiler/internal/bytecode"
)

const StackSize = 2048

// VM is a virtual machine that executes bytecode
type VM struct {
	constants    []string
	instructions bytecode.Instructions
	stack        []string
	sp           int // Stack pointer - points to the next free slot
}

// New creates a new VM instance
func New(bc *bytecode.Bytecode) *VM {
	return &VM{
		constants:    bc.Constants,
		instructions: bc.Instructions,
		stack:        make([]string, StackSize),
		sp:           0,
	}
}

// Run executes the bytecode instructions
func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := bytecode.Opcode(vm.instructions[ip])

		switch op {
		case bytecode.OpConstant:
			// Read the 2-byte operand (constant index)
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1 : ip+3])
			ip += 2 // Skip the operand bytes

			// Push the constant onto the stack
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case bytecode.OpPrint:
			// Pop the value from the stack and print it
			value, err := vm.pop()
			if err != nil {
				return err
			}
			fmt.Println(value)

		case bytecode.OpHalt:
			return nil
		}
	}

	return nil
}

// push pushes a value onto the stack
func (vm *VM) push(value string) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = value
	vm.sp++

	return nil
}

// pop pops a value from the stack
func (vm *VM) pop() (string, error) {
	if vm.sp == 0 {
		return "", fmt.Errorf("stack underflow")
	}

	vm.sp--
	return vm.stack[vm.sp], nil
}

// StackTop returns the top value on the stack without popping it
func (vm *VM) StackTop() (string, error) {
	if vm.sp == 0 {
		return "", fmt.Errorf("empty stack")
	}

	return vm.stack[vm.sp-1], nil
}

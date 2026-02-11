package bytecode

import (
	"encoding/binary"
	"encoding/json"
	"os"
)

// Opcode represents a bytecode instruction
type Opcode byte

const (
	OpConstant Opcode = iota // Load a constant onto the stack
	OpPrint                  // Print the top value from the stack
	OpHalt                   // Halt execution
)

// Instructions is a sequence of bytecode instructions
type Instructions []byte

// Bytecode represents compiled bytecode with its constants
type Bytecode struct {
	Instructions Instructions
	Constants    []string
}

// BytecodeFile represents the serialized format of bytecode
type BytecodeFile struct {
	Version      uint32   `json:"version"`
	Instructions []byte   `json:"instructions"`
	Constants    []string `json:"constants"`
}

// String returns a human-readable representation of an opcode
func (op Opcode) String() string {
	switch op {
	case OpConstant:
		return "OpConstant"
	case OpPrint:
		return "OpPrint"
	case OpHalt:
		return "OpHalt"
	default:
		return "OpUnknown"
	}
}

// WriteToFile writes the bytecode to a file
func (bc *Bytecode) WriteToFile(filename string) error {
	bcFile := BytecodeFile{
		Version:      1,
		Instructions: bc.Instructions,
		Constants:    bc.Constants,
	}

	data, err := json.Marshal(bcFile)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// WriteToBinaryFile writes the bytecode to a binary file
func (bc *Bytecode) WriteToBinaryFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write version (4 bytes)
	if err := binary.Write(f, binary.BigEndian, uint32(1)); err != nil {
		return err
	}

	// Write number of constants (4 bytes)
	if err := binary.Write(f, binary.BigEndian, uint32(len(bc.Constants))); err != nil {
		return err
	}

	// Write each constant (length + string)
	for _, constant := range bc.Constants {
		if err := binary.Write(f, binary.BigEndian, uint32(len(constant))); err != nil {
			return err
		}
		if _, err := f.Write([]byte(constant)); err != nil {
			return err
		}
	}

	// Write instructions length (4 bytes)
	if err := binary.Write(f, binary.BigEndian, uint32(len(bc.Instructions))); err != nil {
		return err
	}

	// Write instructions
	_, err = f.Write(bc.Instructions)
	return err
}

// ReadFromFile reads bytecode from a JSON file
func ReadFromFile(filename string) (*Bytecode, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bcFile BytecodeFile
	if err := json.Unmarshal(data, &bcFile); err != nil {
		return nil, err
	}

	return &Bytecode{
		Instructions: bcFile.Instructions,
		Constants:    bcFile.Constants,
	}, nil
}

// ReadFromBinaryFile reads bytecode from a binary file
func ReadFromBinaryFile(filename string) (*Bytecode, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read version
	var version uint32
	if err := binary.Read(f, binary.BigEndian, &version); err != nil {
		return nil, err
	}

	// Read number of constants
	var numConstants uint32
	if err := binary.Read(f, binary.BigEndian, &numConstants); err != nil {
		return nil, err
	}

	// Read constants
	constants := make([]string, numConstants)
	for i := uint32(0); i < numConstants; i++ {
		var length uint32
		if err := binary.Read(f, binary.BigEndian, &length); err != nil {
			return nil, err
		}

		buf := make([]byte, length)
		if _, err := f.Read(buf); err != nil {
			return nil, err
		}
		constants[i] = string(buf)
	}

	// Read instructions length
	var insLength uint32
	if err := binary.Read(f, binary.BigEndian, &insLength); err != nil {
		return nil, err
	}

	// Read instructions
	instructions := make([]byte, insLength)
	if _, err := f.Read(instructions); err != nil {
		return nil, err
	}

	return &Bytecode{
		Instructions: instructions,
		Constants:    constants,
	}, nil
}

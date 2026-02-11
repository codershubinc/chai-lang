package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"myLang/compiler/internal/compiler"
	laxer "myLang/compiler/internal/lexer"
	"myLang/compiler/internal/parser"
)

const executableTemplate = `package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

var bytecodeData = []byte{ {{.BytecodeHex}} }

const StackSize = 2048

type VM struct {
	constants    []string
	instructions []byte
	stack        []string
	sp           int
}

func NewVM(instructions []byte, constants []string) *VM {
	return &VM{
		constants:    constants,
		instructions: instructions,
		stack:        make([]string, StackSize),
		sp:           0,
	}
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := vm.instructions[ip]
		
		switch op {
		case 0:
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1 : ip+3])
			ip += 2
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		case 1:
			value, err := vm.pop()
			if err != nil {
				return err
			}
			fmt.Println(value)
		case 2:
			return nil
		}
	}
	return nil
}

func (vm *VM) push(value string) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = value
	vm.sp++
	return nil
}

func (vm *VM) pop() (string, error) {
	if vm.sp == 0 {
		return "", fmt.Errorf("stack underflow")
	}
	vm.sp--
	return vm.stack[vm.sp], nil
}

func main() {
	var bcData struct {
		Instructions []byte   ` + "`json:\"instructions\"`" + `
		Constants    []string ` + "`json:\"constants\"`" + `
	}
	
	err := json.Unmarshal(bytecodeData, &bcData)
	if err != nil {
		fmt.Printf("Error decoding bytecode: %v\n", err)
		os.Exit(1)
	}
	
	vm := NewVM(bcData.Instructions, bcData.Constants)
	err = vm.Run()
	if err != nil {
		fmt.Printf("Runtime error: %v\n", err)
		os.Exit(1)
	}
}
`

type TemplateData struct {
	BytecodeHex string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: build <source.chi> [output-name]")
		return
	}

	filename := os.Args[1]
	if len(filename) < 5 || filename[len(filename)-4:] != ".chi" {
		fmt.Println("Error: File extension must be .chi")
		return
	}

	outputName := strings.TrimSuffix(filepath.Base(filename), ".chi")
	if len(os.Args) >= 3 {
		outputName = os.Args[2]
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	code := string(bytes)

	fmt.Println("Compiling", filename, "...")
	l := laxer.NewLexer(code)
	p := parser.NewParser(l)
	program := p.ParseProgram()

	comp := compiler.New()
	err = comp.Compile(program)
	if err != nil {
		fmt.Printf("Compilation error: %v\n", err)
		return
	}

	bc := comp.Bytecode()

	bcData := map[string]interface{}{
		"instructions": bc.Instructions,
		"constants":    bc.Constants,
	}

	jsonData, err := json.Marshal(bcData)
	if err != nil {
		fmt.Printf("Error serializing bytecode: %v\n", err)
		return
	}

	hexParts := make([]string, len(jsonData))
	for i, b := range jsonData {
		hexParts[i] = fmt.Sprintf("0x%02x", b)
	}
	bytecodeHex := strings.Join(hexParts, ", ")

	tmpl, err := template.New("executable").Parse(executableTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	tmpDir, err := os.MkdirTemp("", "chai-build-*")
	if err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	tmpFile := filepath.Join(tmpDir, "main.go")
	f, err := os.Create(tmpFile)
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}

	data := TemplateData{
		BytecodeHex: bytecodeHex,
	}

	err = tmpl.Execute(f, data)
	f.Close()
	if err != nil {
		fmt.Printf("Error generating source: %v\n", err)
		return
	}

	fmt.Println("Building executable", outputName, "...")
	cmd := exec.Command("go", "build", "-o", outputName, tmpFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error building executable: %v\n", err)
		return
	}

	fmt.Printf("✓ Successfully created executable: %s\n", outputName)
	fmt.Println("Run with:", "./"+outputName)
}

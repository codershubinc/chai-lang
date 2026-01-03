# Evaluator Documentation (`compiler/internal/evaluator/eval.go`)

The Evaluator is the "Runtime". It takes the AST and actually executes the instructions. This is an **Interpreter** (it runs the code directly) rather than a Compiler (which would output machine code).

## 1. `Eval(node ast.Node)`

This is a single, recursive function that handles everything.

### How it works

It uses a Go **Type Switch** (`switch node := node.(type)`) to determine exactly what kind of AST node it is dealing with.

### Case 1: `*ast.Program`

- **Input**: The root of the tree.
- **Action**: The program contains a list of statements. The evaluator loops through `node.Statements` and calls `Eval(stmt)` for each one.
- **Logic**: This effectively runs the program line-by-line.

### Case 2: `*ast.ChaiSayStatement`

- **Input**: A specific print command.
- **Action**: It takes the `Value` stored in the node (e.g., "Hello World") and passes it to Go's standard `fmt.Println` function.
- **Result**: The text appears on the user's screen.

## Future Expansion

As we add more features, this file will grow:

- **Math**: We will add cases for `*ast.InfixExpression` (e.g., `5 + 5`) to perform addition.
- **Conditionals**: We will add cases for `*ast.IfExpression` to decide whether to run a block of code.
- **Variables**: We will need to pass an `Environment` object to `Eval` to keep track of variable values.

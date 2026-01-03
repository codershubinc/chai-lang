# AST Documentation (`compiler/internal/ast/ast.go`)

The **Abstract Syntax Tree (AST)** is a tree representation of the source code. Unlike the flat list of tokens from the Lexer, the AST shows the *structure* and *hierarchy* of the code.

## 1. The `Node` Interface
```go
type Node interface {
    TokenLiteral() string
}
```
*   **Purpose**: The base interface. Every single piece of the AST (programs, statements, expressions) is a `Node`.
*   **Method**: `TokenLiteral()` is used mainly for debugging and testing to see the original text.

## 2. The `Statement` Interface
```go
type Statement interface {
    Node
    statementNode()
}
```
*   **Purpose**: Represents a command that performs an action (like `print`, `if`, `return`).
*   **Dummy Method**: `statementNode()` does nothing. It exists only to distinguish `Statement` types from other types (like Expressions) in Go's type system.

## 3. `Program` Struct (The Root)
```go
type Program struct {
    Statements []Statement
}
```
*   **Purpose**: Represents the entire file.
*   **Structure**: It is simply a list (slice) of Statements. When we parse a file, we just add every statement we find to this list.

## 4. `ChaiSayStatement` Struct
```go
type ChaiSayStatement struct {
    Token tokens.Token
    Value string
}
```
*   **Purpose**: Represents the specific command `chai_say "message"`.
*   **Fields**:
    *   `Token`: The `TOKEN_CHAI_SAY` token. We keep this so we know where in the file this statement came from (useful for error messages).
    *   `Value`: The actual string content to be printed (e.g., "Hello").

### Visualizing the Tree
For code:
```plaintext
chai_say "Hello"
chai_say "World"
```

The AST looks like:
```text
Program
├── Statements[0]: ChaiSayStatement
│   └── Value: "Hello"
└── Statements[1]: ChaiSayStatement
    └── Value: "World"
```

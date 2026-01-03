# Parser Documentation (`compiler/internal/parser/parser.go`)

The Parser is the "Grammarian". It takes the flat list of tokens from the Lexer and builds the structured AST defined in `ast.go`.

## 1. The `Parser` Struct

```go
type Parser struct {
    l         *laxer.Lexer
    curToken  tokens.Token
    peekToken tokens.Token
}
```

- **`l`**: A pointer to the Lexer instance.
- **`curToken`**: The token we are currently analyzing.
- **`peekToken`**: The _next_ token.
- **Why two tokens?**: We often need to look ahead. For example, if we see `chai_say`, we need to know if the _next_ token is a string to decide if it's valid.

## 2. `NewParser(l *laxer.Lexer) *Parser`

- **Initialization**: It calls `nextToken()` **twice**.
  1.  First call: Sets `peekToken` to the first token.
  2.  Second call: Moves `peekToken` to `curToken`, and loads the second token into `peekToken`.
  - Result: `curToken` is the first token, `peekToken` is the second.

## 3. `ParseProgram() *ast.Program`

The main loop.

1.  Creates a new `ast.Program`.
2.  Loops until `curToken` is `TOKEN_EOF`.
3.  Inside the loop, calls `parseStatement()`.
4.  If a valid statement is returned, it appends it to `program.Statements`.
5.  Calls `nextToken()` to advance.

## 4. `parseStatement() ast.Statement`

The dispatcher.

- It looks at `p.curToken.Type`.
- If it is `TOKEN_CHAI_SAY`, it calls `parseChaiSayStatement()`.
- If it's anything else, it returns `nil` (ignores it).

## 5. `parseChaiSayStatement() *ast.ChaiSayStatement`

Handles the grammar rule: `chai_say <STRING>`

1.  **Create Node**: Starts building an `ast.ChaiSayStatement` using the current token (`chai_say`).
2.  **Validation (Peek)**: Checks `p.peekToken.Type`.
    - It **expects** the next token to be `TOKEN_STRING`.
    - If it is NOT a string, it returns `nil` (syntax error).
3.  **Advance**: Calls `nextToken()` to move focus to the string token.
4.  **Extract Value**: Sets `stmt.Value` to the literal value of the string token.
5.  **Return**: Returns the completed AST node.

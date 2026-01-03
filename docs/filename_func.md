# Project Function Reference (`filename_func.md`)

This document provides a detailed breakdown of every function in the project, organized by the file they belong to. It explains what each function does, how it works, and where it is used.

---

## 1. File: `compiler/cmd/main.go`

This is the entry point of the application.

### Function: `main()`

- **Signature**: `func main()`
- **Purpose**: The "Conductor" of the compiler. It orchestrates the entire flow from reading the file to executing the code.
- **Working**:
  1.  **Validation**: Checks if a command-line argument is provided and if it ends with `.chi`.
  2.  **File Reading**: Uses `os.ReadFile` to load the source code into a string.
  3.  **Pipeline Execution**:
      - Initializes the **Lexer** with the source code.
      - Initializes the **Parser** with the Lexer.
      - Calls `ParseProgram()` to generate the AST.
      - Calls `evaluator.Eval()` to run the AST.
- **Use**: Automatically called by the Go runtime when you run the program.

---

## 2. File: `compiler/internal/laxer/laxer.go`

The Lexer (or Tokenizer) is the first phase of the compiler. Its job is to read the raw source code (a long string of characters) and group them into meaningful units called **Tokens**. It's like reading a sentence and identifying "Subject", "Verb", "Object".

### Struct: `Lexer`

- **`input`**: The complete source code string (e.g., `chai_say "hello"`).
- **`position`**: The index of the character we are _currently_ inspecting.
- **`readPosition`**: The index of the _next_ character (used to peek ahead).
- **`ch`**: The actual byte value of the character at `position`.

### Function: `NewLexer(input string) *Lexer`

- **Purpose**: Initializes a new Lexer machine.
- **Working**:
  1.  Creates a `Lexer` struct with the provided `input`.
  2.  Calls `readChar()` immediately. This is crucial because we need to load the very first character into `l.ch` before we can start analyzing.
- **Use**: Called in `main.go` to start the pipeline.

### Function: `readChar()`

- **Purpose**: The engine that moves the Lexer forward. It advances the "cursor" by one character.
- **Working**:
  1.  **Check Bounds**: It checks if `readPosition` has gone past the end of the input string.
  2.  **Handle EOF**: If we are at the end, it sets `l.ch = 0`. In ASCII, 0 is the NUL character, which we use to signify "End of File".
  3.  **Advance**: If not at the end, it sets `l.ch` to the character at `readPosition`.
  4.  **Update Pointers**: It updates `l.position` to be the current `l.readPosition`, and increments `l.readPosition` by 1.
- **Use**: Called internally by almost every other function (`NextToken`, `readString`, `readIdentifier`) whenever they need to consume a character.

### Function: `NextToken() tokens.Token`

- **Purpose**: The brain of the Lexer. It looks at the current character and decides what Token to produce.
- **Working**:
  1.  **Skip Noise**: Calls `skipWhitespace()` to eat up spaces, tabs, and newlines. We don't care about indentation in this language.
  2.  **Analyze Character**: It switches on `l.ch`:
      - **Case `"`**: It knows a string is starting. It calls `readString()` to get the full text and returns a `TOKEN_STRING`.
      - **Case `0` (EOF)**: It returns a `TOKEN_EOF` to tell the parser "we are done".
      - **Default (Letters)**: If `isLetter(l.ch)` is true, it knows a word is starting.
        - It calls `readIdentifier()` to read the whole word (e.g., "chai_say").
        - It checks if this word is a known keyword (like "chai_say"). If yes, it returns `TOKEN_CHAI_SAY`. If no, it would return `TOKEN_IDENTIFIER` (though we haven't implemented variables yet).
      - **Default (Unknown)**: If it's none of the above, it returns `TOKEN_ILLEGAL`.
  3.  **Advance**: After creating the token, it calls `readChar()` to move past it (except for identifiers/strings which handle their own advancement).
- **Use**: The **Parser** calls this repeatedly. "Give me the next token. Okay, give me the next one."

### Function: `readString() string`

- **Purpose**: Extracts text inside quotes.
- **Working**:
  1.  Remembers the `position + 1` (where the text starts, skipping the opening quote).
  2.  Enters a loop: calls `readChar()` repeatedly.
  3.  **Stop Condition**: The loop breaks if it hits a closing quote `"` or EOF `0`.
  4.  Returns the slice of the input string from the start position to the current position.
- **Use**: Called by `NextToken` when it encounters a `"`.

### Function: `readIdentifier() string`

- **Purpose**: Extracts a word (keyword or variable name).
- **Working**:
  1.  Remembers the start `position`.
  2.  Enters a loop: calls `readChar()` as long as `isLetter(l.ch)` is true.
  3.  Returns the slice of the input string from the start to the end.
- **Use**: Called by `NextToken` when it encounters a letter.

### Function: `skipWhitespace()`

- **Purpose**: Cleans up the input stream.
- **Working**:
  - Checks if `l.ch` is a space `' '`, tab `'\t'`, newline `'\n'`, or return `'\r'`.
  - If yes, calls `readChar()` to ignore it and checks again.
  - Repeats until a non-whitespace character is found.
- **Use**: Called at the very start of `NextToken`.

### Function: `isLetter(ch byte) bool`

- **Purpose**: Defines the alphabet of our identifiers.
- **Working**:
  - Returns `true` if `ch` is between 'a' and 'z'.
  - Returns `true` if `ch` is between 'A' and 'Z'.
  - **Crucial**: Returns `true` if `ch` is `_` (underscore). This allows `chai_say` to be read as one word. Without this, it would be read as `chai` (word) and `_` (illegal).
- **Use**: Used by `readIdentifier` to know when a word ends.

---

## 3. File: `compiler/internal/parser/parser.go`

The Parser turns Tokens into an Abstract Syntax Tree (AST).

### Function: `NewParser(l *laxer.Lexer) *Parser`

- **Purpose**: Constructor. Creates a new Parser.
- **Working**: Stores the Lexer and calls `nextToken()` **twice** to fill `curToken` (current) and `peekToken` (next) for lookahead.
- **Use**: Called in `main.go`.

### Function: `nextToken()`

- **Purpose**: Advances the token stream.
- **Working**: Moves `peekToken` to `curToken`, and asks the Lexer for a new `peekToken`.
- **Use**: Called internally whenever we consume a token.

### Function: `ParseProgram() *ast.Program`

- **Purpose**: The main parsing loop.
- **Working**: Creates a root `Program` node. Loops until `TOKEN_EOF`, calling `parseStatement()` repeatedly and adding valid statements to the Program's list.
- **Use**: Called in `main.go`.

### Function: `parseStatement() ast.Statement`

- **Purpose**: Decides how to parse the current statement based on the token type.
- **Working**: Checks `p.curToken.Type`. If it is `TOKEN_CHAI_SAY`, it delegates to `parseChaiSayStatement`. Otherwise returns `nil`.
- **Use**: Called by `ParseProgram`.

### Function: `parseChaiSayStatement() *ast.ChaiSayStatement`

- **Purpose**: Handles the specific grammar for `chai_say`.
- **Working**:
  1.  Creates a `ChaiSayStatement` node.
  2.  Checks if the _next_ token (`peekToken`) is a `TOKEN_STRING`.
  3.  If yes, advances tokens and saves the string value.
  4.  Returns the populated node.
- **Use**: Called by `parseStatement`.

---

## 4. File: `compiler/internal/evaluator/eval.go`

The Evaluator executes the AST.

### Function: `Eval(node ast.Node)`

- **Purpose**: Recursively executes the code represented by the AST.
- **Working**: Uses a type switch to determine what kind of node it is processing:
  - `*ast.Program`: Loops through its `Statements` list and calls `Eval` on each.
  - `*ast.ChaiSayStatement`: Prints the node's `Value` to the console using `fmt.Println`.
- **Use**: Called in `main.go` to run the program, and recursively calls itself.

---

## 5. File: `compiler/internal/ast/ast.go`

The AST (Abstract Syntax Tree) is the structural representation of the code. While the Lexer deals with a flat list of tokens, the AST deals with hierarchy and meaning.

### Interface: `Node`

- **Purpose**: The base interface for _everything_ in our tree. Whether it's a whole program, a statement, or an expression, it is a `Node`.
- **Method**: `TokenLiteral() string` - A helper method that lets us retrieve the original text of the token associated with this node (mostly for debugging).

### Interface: `Statement`

- **Purpose**: Represents a line of code that performs an action but doesn't produce a value (e.g., `chai_say "hi"` is a statement. `5 + 5` is an expression).
- **Inheritance**: It embeds `Node`, so every `Statement` is also a `Node`.
- **Method**: `statementNode()` - This is a "dummy" method. It doesn't do anything logic-wise. Its only purpose is to distinguish `Statement` types from `Expression` types in Go's type system. Only structs that implement this method can be used where a `Statement` is expected.

### Struct: `Program`

- **Purpose**: The root node of the entire tree. It represents the whole file.
- **Field**: `Statements []Statement` - A list (slice) that holds every top-level statement found in the file.
- **Function**: `TokenLiteral()`
  - **Working**: If the program has statements, it returns the literal of the first one. Otherwise, returns empty string.
  - **Use**: Debugging.

### Struct: `ChaiSayStatement`

- **Purpose**: Represents the specific command `chai_say "..."`.
- **Fields**:
  - `Token`: The actual `TOKEN_CHAI_SAY` token (useful for error reporting, knowing line numbers, etc.).
  - `Value`: The string content that comes after the keyword (e.g., "hello").
- **Function**: `statementNode()`
  - **Working**: Empty. Just satisfies the interface.
- **Function**: `TokenLiteral()`
  - **Working**: Returns `cs.Token.Literal` (which would be "chai_say").

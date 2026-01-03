# Lexer Documentation (`compiler/internal/laxer/laxer.go`)

The Lexer (or Tokenizer) is the first phase of the compiler pipeline. It reads source code character-by-character and groups them into Tokens.

## 1. The `Lexer` Struct
```go
type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           byte
}
```
*   **`input`**: The entire source code string.
*   **`position`**: The index of the character we are *currently* looking at.
*   **`readPosition`**: The index of the *next* character. We need this to "peek" ahead (e.g., to see if `!` is followed by `=`).
*   **`ch`**: The actual character (byte) at `position`.

## 2. `NewLexer(input string) *Lexer`
Initializes the lexer.
*   **Crucial Step**: It calls `readChar()` immediately. This loads the very first character of the file into `l.ch` so the lexer is ready to start.

## 3. `readChar()`
The engine of the lexer. It advances the cursor.
*   **Logic**:
    1.  Checks if `readPosition` is past the end of the input.
    2.  If yes, sets `ch` to `0` (NUL character), signifying EOF.
    3.  If no, sets `ch` to the next character.
    4.  Updates `position` and `readPosition`.

## 4. `NextToken() tokens.Token`
The main brain. It looks at `l.ch` and decides what token to create.

*   **Step 1: Skip Whitespace**: Calls `skipWhitespace()` to ignore spaces, tabs, and newlines.
*   **Step 2: Identify Character**:
    *   `"`: Starts a string. Calls `readString()`.
    *   `0`: End of file. Returns `TOKEN_EOF`.
    *   **Letters**: Starts an identifier or keyword. Calls `readIdentifier()`.
        *   **Keyword Check**: After reading the word, it checks if it matches `chai_say`. If so, returns `TOKEN_CHAI_SAY`. Otherwise, returns `TOKEN_ILLEGAL` (since we don't support variables yet).
*   **Step 3: Advance**: Calls `readChar()` to prepare for the next call.

## 5. Helper Functions

### `readString() string`
*   **Goal**: Extract text between quotes.
*   **How**: Loops `readChar()` until it hits another `"` or EOF. Returns the text *inside* the quotes (excluding the quotes themselves).

### `readIdentifier() string`
*   **Goal**: Extract a word.
*   **How**: Loops `readChar()` as long as `isLetter()` returns true.

### `isLetter(ch byte) bool`
*   **Goal**: Define what a "word" character is.
*   **Logic**: Returns true for `a-z`, `A-Z`, and `_`.
*   **Importance**: The `_` is critical. Without it, `chai_say` would be read as `chai` (word) and `_` (symbol), breaking the parser.

### `skipWhitespace()`
*   **Goal**: Ignore formatting.
*   **How**: Loops while `ch` is space, tab, newline, or return.

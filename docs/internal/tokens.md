# Tokens Documentation (`compiler/internal/tokens/token.go`)

This file defines the fundamental building blocks of the Chai language: the **Tokens**.

## What is a Token?

A token is the smallest meaningful unit of code. Just as a sentence is made of words and punctuation, code is made of tokens. The Lexer's job is to read raw text and produce these tokens.

## 1. `TokenType` (Type Definition)

```go
type TokenType string
```

We use a string alias for token types. This makes debugging easier because we can print the type directly (e.g., seeing "CHAI_SAY" is better than seeing a number like `4`).

## 2. Constants (The Vocabulary)

We define constants for every type of token our language supports.

- **`TOKEN_EOF` ("EOF")**:
  - Stands for "End Of File".
  - **Purpose**: Tells the parser "There is no more code to read." It prevents infinite loops.
- **`TOKEN_CHAI_SAY` ("CHAI_SAY")**:
  - **Purpose**: Represents the keyword `chai_say`.
  - **Usage**: When the lexer sees the text `chai_say`, it generates this token.
- **`TOKEN_STRING` ("STRING")**:
  - **Purpose**: Represents text data enclosed in quotes.
  - **Usage**: When the lexer sees `"hello"`, it generates a STRING token with the literal value `hello`.
- **`TOKEN_ILLEGAL` ("ILLEGAL")**:
  - **Purpose**: Represents something the lexer doesn't understand.
  - **Usage**: If you type a character like `@` or `#` (assuming they aren't allowed), the lexer marks it as ILLEGAL.

## 3. `Token` Struct

```go
type Token struct {
    Type    TokenType
    Literal string
}
```

This is the data structure passed around the compiler.

- **`Type`**: The category of the token (e.g., `TOKEN_STRING`).
- **`Literal`**: The actual text found in the source code (e.g., `"Hello World"`).

### Example

For the code: `chai_say "Hi"`

1.  **First Token**:
    - Type: `TOKEN_CHAI_SAY`
    - Literal: `chai_say`
2.  **Second Token**:
    - Type: `TOKEN_STRING`
    - Literal: `Hi`

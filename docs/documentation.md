# MyLang (Chai) Compiler Documentation

Welcome! This documentation is designed for beginners who want to understand how a programming language is built from scratch. We will explore the "Chai" language compiler we've built, explaining every concept and file in detail.

## 1. The Big Picture: How a Compiler Works

> **Visual Guide**: Check out [docs/diagrams.md](diagrams.md) for flowcharts and diagrams of this process!

Imagine reading a book. To understand a sentence like "The cat sleeps.", your brain goes through several steps. A compiler does something very similar with code.

Our compiler follows a **4-step pipeline**:

1.  **Lexical Analysis (Lexing)**:
    *   *Analogy*: Identifying individual words and punctuation.
    *   *Input*: `chai_say "Hello"`
    *   *Output*: `[TOKEN_CHAI_SAY, TOKEN_STRING("Hello")]`
    *   *What it does*: It reads the raw text (source code) character by character and groups them into meaningful units called **Tokens**.

2.  **Parsing**:
    *   *Analogy*: Understanding the grammar. Knowing that "The cat sleeps" is a valid sentence (Subject + Verb), but "Sleeps cat the" is not.
    *   *Input*: Tokens from the Lexer.
    *   *Output*: **Abstract Syntax Tree (AST)**.
    *   *What it does*: It checks if the tokens follow the rules of the language and builds a structured tree representation of the code.

3.  **Abstract Syntax Tree (AST)**:
    *   *Analogy*: A diagram of the sentence structure.
    *   *What it does*: It is a data structure that represents the logical structure of the program. It removes unnecessary details (like whitespace) and keeps the logic.

4.  **Evaluation**:
    *   *Analogy*: Understanding the meaning and taking action (imagining the sleeping cat).
    *   *Input*: The AST.
    *   *Output*: The actual result (printing text to the screen).
    *   *What it does*: It walks through the AST and executes the instructions found there.

---

## 2. Project Structure

Here is how our project is organized:

```text
myLang/
├── compiler/
│   ├── cmd/
│   │   └── main.go           # The Entry Point (The Conductor)
│   └── internal/
│       ├── tokens/
│       │   └── token.go      # The Definitions (The Alphabet)
│       ├── laxer/
│       │   └── laxer.go      # The Lexer (The Reader)
│       ├── ast/
│       │   └── ast.go        # The AST (The Structure)
│       ├── parser/
│       │   └── parser.go     # The Parser (The Grammarian)
│       └── evaluator/
│           └── eval.go       # The Evaluator (The Executor)
├── test.chi                  # Our source code file
└── go.mod                    # Go module definition
```

---

## 3. Deep Dive into the Code

Let's look at each part of the machine.

### A. `tokens/token.go` (The Alphabet)

Before we can read words, we need to define what "words" exist in our language.

*   **`TokenType`**: A string that represents the category of the token (e.g., "STRING", "CHAI_SAY").
*   **`Token` struct**: This holds the actual piece of data.
    *   `Type`: What kind of token is this?
    *   `Literal`: What was the actual text written? (e.g., "Hello").

**Key Tokens defined:**
*   `TOKEN_CHAI_SAY`: The keyword `chai_say`.
*   `TOKEN_STRING`: Anything inside quotes `"like this"`.
*   `TOKEN_EOF`: "End Of File" - tells us we are done reading.

### B. `laxer/laxer.go` (The Reader)

The Lexer is the "dumb" part of the compiler. It doesn't understand code; it just recognizes words.

*   **`Lexer` struct**: Keeps track of:
    *   `input`: The whole source code string.
    *   `position`: Where we are currently looking.
    *   `ch`: The current character we are looking at.

*   **`readChar()`**: Moves our cursor forward by one character.
*   **`NextToken()`**: The most important function!
    1.  It skips whitespace (spaces, tabs).
    2.  It looks at the current character (`l.ch`).
    3.  If it sees a quote `"`, it calls `readString()` to grab the whole text inside.
    4.  If it sees a letter, it calls `readIdentifier()` to grab the whole word (like `chai_say`).
    5.  **Crucial Logic**: It checks if the word is a known keyword. If it sees `chai_say`, it creates a `TOKEN_CHAI_SAY`.

*   **The `isLetter` Fix**:
    *   We define a "letter" as `a-z`, `A-Z`, or `_`.
    *   *Why?* Because our keyword `chai_say` has an underscore. If we didn't include `_`, the lexer would stop at `chai` and get confused.

### C. `ast/ast.go` (The Structure)

This file defines the "shapes" of our code.

*   **`Node` Interface**: Everything in our tree is a Node.
*   **`Program` struct**: The root of the tree. It holds a list of `Statements`.
*   **`ChaiSayStatement` struct**: Represents a specific command.
    *   It holds the `Token` (for debugging) and the `Value` (the message to print).

### D. `parser/parser.go` (The Grammarian)

The Parser is the "smart" part. It asks the Lexer for tokens and builds the AST.

*   **`Parser` struct**: Holds the Lexer and keeps track of the `curToken` (current) and `peekToken` (next). We look ahead to make decisions.
*   **`ParseProgram()`**:
    *   Loops through every token until it hits `EOF`.
    *   Calls `parseStatement()` for each one.
*   **`parseStatement()`**:
    *   Checks the current token type.
    *   If it is `TOKEN_CHAI_SAY`, it calls `parseChaiSayStatement()`.
*   **`parseChaiSayStatement()`**:
    *   It expects the *current* token to be `chai_say`.
    *   It expects the *next* token (`peekToken`) to be a `STRING`.
    *   If both are true, it creates a `ChaiSayStatement` node and saves the string value into it.

### E. `evaluator/eval.go` (The Executor)

This is where the magic happens. The code actually runs here.

*   **`Eval(node)`**: A recursive function.
    *   If the node is a `Program`, it loops through all statements and calls `Eval` on them.
    *   If the node is a `ChaiSayStatement`, it takes the `Value` and uses Go's `fmt.Println` to print it to your terminal.

### F. `cmd/main.go` (The Conductor)

This ties everything together.

1.  **Read File**: It checks if you provided a `.chi` file and reads the text.
2.  **Setup Pipeline**:
    *   `l := laxer.NewLexer(code)`
    *   `p := parser.NewParser(l)`
    *   `program := p.ParseProgram()`
3.  **Execute**:
    *   `evaluator.Eval(program)`

---

## 4. The "Chai" Language Syntax

Currently, our language is very simple. It supports one command:

```plaintext
chai_say "Your Message Here"
```

*   **Keywords**: `chai_say` (prints to console).
*   **Data Types**: Strings (text inside double quotes).

---

## 5. How to Run It

1.  **Write Code**: Create a file named `test.chi` with the following content:
    ```plaintext
    chai_say "Hello World"
    ```

2.  **Run Compiler**: Open your terminal in the project folder and run:
    ```bash
    go run ./compiler/cmd/main.go test.chi
    ```

3.  **See Output**:
    ```text
    Hello World
    ```

---

## Summary

You have built a working interpreter!
1.  **Lexer** broke `chai_say "Hi"` into `[CHAI_SAY, "Hi"]`.
2.  **Parser** saw `CHAI_SAY` and looked for a string, creating a `ChaiSayStatement`.
3.  **Evaluator** saw the `ChaiSayStatement` and ran `fmt.Println("Hi")`.

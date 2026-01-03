# Visualizing the Compiler (Mermaid Diagrams)

This document uses Mermaid diagrams to visualize how the different parts of the compiler interact and work internally.

## 1. High-Level Architecture

This diagram shows the data flow from your source code file to the final output.

```mermaid
flowchart LR
    Source["Source Code (.chi)"] -->|Input| Lexer
    Lexer -->|Generates| Tokens
    Tokens -->|Input| Parser
    Parser -->|Builds| AST[Abstract Syntax Tree]
    AST -->|Input| Evaluator
    Evaluator -->|Executes| Output[Terminal Output]

    style Source fill:#f9f,stroke:#333,stroke-width:2px
    style AST fill:#bbf,stroke:#333,stroke-width:2px
    style Output fill:#bfb,stroke:#333,stroke-width:2px
```

---

## 2. The Lexer Process (`NextToken`)

How the Lexer decides what token to create when it reads a character.

```mermaid
flowchart TD
    Start([NextToken Called]) --> SkipWS[Skip Whitespace]
    SkipWS --> CheckChar{Check Current Char}

    CheckChar -->|Quote| ReadString[Read String Literal]
    ReadString --> MakeString[Return TOKEN_STRING]

    CheckChar -->|Letter| ReadIdent[Read Identifier/Word]
    ReadIdent --> CheckKey{"Is it 'chai_say'?"}
    CheckKey -- Yes --> MakeKey[Return TOKEN_CHAI_SAY]
    CheckKey -- No --> MakeIllegal[Return TOKEN_ILLEGAL]

    CheckChar -->|"EOF (0)"| MakeEOF[Return TOKEN_EOF]

    CheckChar -- Other --> MakeIll[Return TOKEN_ILLEGAL]

    MakeString --> Finish([Return Token])
    MakeKey --> Finish
    MakeEOF --> Finish
    MakeIllegal --> Finish
    MakeIll --> Finish
```

---

## 3. The Parser Interaction

How the Parser talks to the Lexer to build the AST.

```mermaid
sequenceDiagram
    participant Main
    participant Parser
    participant Lexer

    Main->>Parser: ParseProgram()

    loop Until EOF
        Parser->>Parser: Check curToken.Type

        alt is TOKEN_CHAI_SAY
            Parser->>Parser: parseChaiSayStatement()
            Parser->>Lexer: Check peekToken (Lookahead)

            alt peekToken is STRING
                Parser->>Lexer: nextToken() (Advance)
                Parser->>Parser: Create ChaiSayStatement Node
                Note right of Parser: Stores "Value" from token
            else peekToken is NOT STRING
                Parser->>Parser: Return nil (Ignore/Error)
            end

        else is Unknown
            Parser->>Parser: Return nil
        end

        Parser->>Lexer: nextToken() (Move to next)
    end

    Parser-->>Main: Return AST (Program)
```

---

## 4. AST Visualization

What the data structure looks like for a simple program.

**Code:**

```plaintext
chai_say "Hello"
chai_say "World"
```

**Tree Structure:**

```mermaid
graph TD
    Root[Program Node]

    Stmt1["Statement 1: ChaiSayStatement"]
    Stmt2["Statement 2: ChaiSayStatement"]

    Root -->|Statements List| Stmt1
    Root -->|Statements List| Stmt2

    Token1[Token: CHAI_SAY]
    Value1["Value: 'Hello'"]

    Token2[Token: CHAI_SAY]
    Value2["Value: 'World'"]

    Stmt1 --> Token1
    Stmt1 --> Value1

    Stmt2 --> Token2
    Stmt2 --> Value2

    style Root fill:#f96,stroke:#333,stroke-width:4px
    style Stmt1 fill:#69f,stroke:#333,stroke-width:2px
    style Stmt2 fill:#69f,stroke:#333,stroke-width:2px
```

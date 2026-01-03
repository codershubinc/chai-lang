# chai-lang

A programming language brewed for tea lovers.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/codershubinc/chai-lang.git
   cd chai-lang
   ```

2. Build the compiler:

   ```bash
   go build -o chai compiler/cmd/main.go
   ```

## Usage

Run a `.chi` file using the compiled binary:

```bash
./chai test.chi
```

## Syntax

### Printing to Console

Currently, Chai uses `chai_say` to print text to the console.

**Example:**

```javascript
chai_say "Hello, World!"
chai_say "Time for a tea break."
```

## Roadmap

We are brainstorming new syntax ideas! Check out the `ideas/` folder for what's brewing (like `sip`, `boil_chai`, and `ingredient`).

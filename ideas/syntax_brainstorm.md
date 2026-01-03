# Chai Language Syntax Brainstorming

This document tracks ideas for the syntax and structure of the **Chai** programming language. The core theme is "Making Tea".

## 1. Program Lifecycle (Start & End)

The program should feel like a recipe or a ritual of making tea.

### Idea A: The "Cooking" Metaphor (User's Suggestion)

- **Start**: `turn_on_gas`
- **End**: `serve_chai`
- **Vibe**: Very procedural and action-oriented.

### Idea B: The "Preparation" Metaphor

- **Start**: `boil_water`
- **End**: `sip_chai`
- **Vibe**: Focuses on the ingredients and the result.

### Idea C: The "Direct" Metaphor

- **Start**: `start_chai`
- **End**: `end_chai`
- **Vibe**: Simple and clear, but less thematic.

---

## 2. Output (Printing to Console)

How do we display text to the user?

- **`sip`**: (e.g., `sip "Delicious"`) - The user "consumes" the output.
- **`spill`**: (e.g., `spill "The truth"`) - Play on "spilling the tea" (gossip).
- **`serve`**: (e.g., `serve "Here is data"`) - You are serving the output.
- **`bol`**: (e.g., `bol "Namaste"`) - Hindi for "speak".

---

## 3. Variable Declaration (Future Ideas)

Defining values could look like adding ingredients.

- **`add_sugar`**: For constants or strings?
  - `add_sugar name = "Swap"`
- **`add_milk`**: For integers?
  - `add_milk age = 25`
- **`mix`**: General variable declaration.
  - `mix x = 10`
- **`ingredient`**: Formal declaration.
  - `ingredient count = 5`

---

## 4. Control Flow (Future Ideas)

### Conditionals (If/Else)

- **`if_hot`** / **`else_cold`**
- **`check_taste`**

### Loops

- **`stir`** (e.g., `stir 10 times`) - Represents the circular motion of mixing.
- **`boil_while`** - Loop while a condition is true.

---

## 5. Example Code Draft

```javascript
turn_on_gas

    // Printing
    sip "Making the tea..."

    // Variables (Hypothetical)
    add_sugar name = "Chai"

    // Loop (Hypothetical)
    stir (5 times) {
        sip "Mixing..."
    }

serve_chai
```

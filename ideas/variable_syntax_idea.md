# Variable Syntax Ideas

Using ingredients to define data types.

## General Syntax

### Option A: Explicit Typing (Thematic)

```javascript
ingredient <type> <name> = <value>;
// Example: ingredient sugar count = 5;
```

### Option B: Simple / Inferred (User Suggestion)

Just use `ingredient` for everything, like `var` or `let`. The type is guessed from the value.

```javascript
ingredient <name> = <value>;
// Example: ingredient name = "Swap";
```

## Data Types (For Option A)

### 1. Integers (Whole Numbers) -> `sugar`

Sugar is added in spoons (discrete units).

- **Syntax**: `ingredient sugar count = 5;`

### 2. Strings (Text) -> `flavor`

Adds character and description to the chai.

- **Syntax**: `ingredient flavor message = "Hello World";`

### 3. Floats (Decimals) -> `milk`

Liquids are measured in volume and can be fractional.

- **Syntax**: `ingredient milk quantity = 1.5;`

### 4. Booleans (True/False) -> `ginger`

It's either in the tea or it isn't.

- **Syntax**: `ingredient ginger is_ready = true;`

### 5. Arrays/Lists -> `elaichi`

Cardamom pods contain multiple seeds, just like arrays contain multiple values.

- **Syntax**: `ingredient elaichi scores = [10, 20, 30];`

### 6. Objects/Structs -> `masala`

A complex mix of different spices.

- **Syntax**:
  ```javascript
  ingredient masala user = {
      name: "Swap",
      age: 25
  };
  ```

### 7. Null/Void -> `water`

The base, or nothingness before the tea is made.

- **Syntax**: `return water;`

## Example Program

```javascript
turn_on_gas

    ingredient sugar spoons = 2;
    ingredient flavor type = "Masala";
    ingredient milk amount = 0.5;

    sip "Making " + type + " chai with " + spoons + " spoons of sugar."

serve_chai
```

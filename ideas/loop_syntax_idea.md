# Loop Syntax Ideas

Brainstorming how loops (iteration) should work in Chai.

## 1. The "Boil" Loop (While/Until Loop)

Represents the time spent boiling the tea until it's ready.

**Syntax:**

```javascript
boil_chai until <condition> {
    // code block
}
```

**Example:**

```javascript
// Loop until temperature reaches 100
boil_chai until temperature == 100 {
    sip "Heating up..."
    temperature = temperature + 1
}
```

## 2. The "Stir" Loop (For Loop / Fixed Iteration)

Represents the repetitive action of stirring the tea. Good for iterating a specific number of times.

**Syntax:**

```javascript
stir <number> times {
    // code block
}
```

**Example:**

```javascript
// Run exactly 5 times
stir 5 times {
    sip "Mixing ingredients..."
}
```

## 3. The "Pour" Loop (ForEach Loop)

Represents pouring tea into multiple cups. Good for iterating over a list or array.

**Syntax:**

```javascript
pour <
  collection >
  into <
  variable >
  {
    // code block
  };
```

**Example:**

```javascript
// Iterate over a list of guests
pour guests into guest {
    sip "Serving " + guest
}
```

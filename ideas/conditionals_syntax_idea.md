# Conditional Syntax Ideas

Brainstorming how to make decisions (If/Else) in Chai.

## 1. The "Taste" Metaphor (Sensory)

Making tea involves tasting to see if it's ready or needs more ingredients.

**Syntax:**

```javascript
taste (sugar > 5) {
    sip "Too sweet!"
} otherwise {
    sip "Perfect."
}
```

## 2. The "Temperature" Metaphor

Tea is either hot (true/good) or cold (false/bad).

**Syntax:**

```javascript
if_hot (is_ready) {
    serve_chai
} else_cold {
    boil_chai
}
```

## 3. The "Check" Metaphor (Procedural)

Simple and clear.

**Syntax:**

```javascript
check (water_level < 10) {
    add_water
}
```

## 4. Standard (Boring but Safe)

```javascript
if (x == 5) {
  // code
} else {
  // code
}
```
## 5. The "Hinglish" Metaphor (Authentic)
Using Hindi words since "Chai" is Hindi.
**Syntax:**
```javascript
agar (sugar == 0) {
    sip "Add sugar!"
} nahi_toh {
    sip "Good to go."
}
```

## 6. The "Experiment" Metaphor
Trying a sip to see if it's good.
**Syntax:**
```javascript
try_sip (is_hot) {
    drink
} spit_out {
    wait
}
```

## 7. The "When" Metaphor (State Based)
Checking the state of the chai.
**Syntax:**
```javascript
when (color == "brown") {
    stop_boiling
} fallback {
    keep_boiling
}
```
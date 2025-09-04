# Dev Journal — First Steps with the Lexer

## 26th of August 2025

## Understanding the Compiler Pipeline

- **Lexer**: turns raw code into tokens.
- **Parser**: consumes tokens and produces an AST.
- Tokens are the **building blocks**; parser uses them to form structure.

## How the Lexer Works

- Reads input **character by character**:
  - Collects letters → check if keyword (`let`, `fn`) or identifier.
  - Collects digits → number literal.
  - Matches single/double-char operators (`=`, `:=`, `==`).
  - Skips whitespace (unless significant).
- Example:
  let x = 5 + 5;
  [LET, IDENT(x), =, INT(5), +, INT(5), ;]

## Tests Are Simple

- **Take input -> compare the output -> to expected output for this data**
- Input → Expected Tokens → Compare. - Go’s testing.T = same idea as manual assertion checks. - Benefit: confidence after refactor, instant feedback. - Example test:
  Input:
  **=+(){},;**
  Expected:
  **_[=, +, (, ), {, }, ,, ;]_**

## 29th of August 2025

### **Struct in Go**

- behaves like class -> you can initialize an instance of it
- afterwards you can use its methods by calling it on struct instance
- classic, is it a part of oop, i think so, but which one, its encapsulation GOAT of the oop isnt it?
- so how do you build a struct?

  ```go
  type StructName struct {
    input string
    field fieldType
    field fieldType...
  }
  ```

- function to initialize the struct

  ```go
  func New(input string) *StructName {
    sn := &StructName{input: input}
    return sn
  }
  ```

- function which is part of the struct inner methods

  ```go
  func (sn *StructName) FunctionName(param paramType) return returnType {
    ret := return.Type{finnish: "success"}
    return ret
  }
  ```

---

### Building the Lexer

- Created a **`Lexer` struct** to hold state:

  - `Input` → the full code string
  - `Position` → current character index
  - `ReadPosition` → next character index
  - `ch` → current character

- Wrote a **constructor**:

  ```go
  func New(input string) *Lexer {
      lexer := &Lexer{Input: input}
      return lexer
  }
  ```

  - works like a class constructor
  - now I can call `lexer := New("let x = 5;")`

- Added method **`NextToken()`**:

  - responsible for producing tokens one by one
  - right now just returns a `token.Token` with `Literal: char`

- Added **`readChar()`**:

  - grabs current character from input string (`l.Input[l.Position]`)
  - stores it in `l.ch`
  - increments both `Position` and `ReadPosition`
  - returns the current character
  - acts like a cursor moving forward

- Added **`peek()`**:

  - looks ahead at the next character (`l.Input[l.ReadPosition]`) without advancing
  - useful for two-character tokens like `==` or `:=`

- Open questions / thoughts:

  - Why does `NextToken` take a char instead of a token? (eventually it should combine chars into identifiers, numbers, etc.)
  - `peek()` could be extended to check multiple characters (`++`, `+=`, etc.)

---

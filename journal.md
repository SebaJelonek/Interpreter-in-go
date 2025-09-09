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

## 5th of September 2025
### **Lexer Implementation Refinement**

- **Centralized `NextToken()` Logic:** The primary change in my understanding is that `NextToken()` should act as the central orchestrator. It's not just for single characters; it's the main method responsible for returning a complete token. It decides what kind of token is coming up and delegates to helper methods for multi-character tokens. 

- **Helper Methods for Multi-Character Tokens:** I realized the need for dedicated helper functions like `readIdentifier()` and `readNumber()`. Their sole purpose is to consume the entire token string—a sequence of letters and/or digits—and return the resulting literal. This keeps `NextToken()` clean and focused.

- **Advancing the Cursor:** A crucial bug fix was understanding that these helper methods **must** advance the lexer's position. My initial attempts failed because the loop would check the same character repeatedly without moving forward. The `readChar()` method needs to be called inside the loop to move the "cursor" through the input string.

- **Identifier vs. Keyword:** I grasped the critical distinction between a generic identifier (like `five` or `result`) and a keyword (`let`, `fn`). My lexer must now include a lookup step after reading an identifier. I'll use a `map[string]token.TokenType` to check if the literal is a reserved keyword; otherwise, it's a standard identifier.

- **`isLetter()` and `isDigit()` Helpers:** I learned that simple, direct boolean checks like `ch >= 'a' && ch <= 'z'` are far cleaner and more readable than complex ASCII range logic with explicit exclusions. This makes the code more robust and easier to maintain.

---

### **Refining My `NextToken()` Logic**

- **`NextToken()` pseudo-code:**
  - Skip any whitespace first.
  - Look at the current character (`l.ch`).
  - If `isLetter(l.ch)` is true:
    - Call `literal := l.readIdentifier()`.
    - Use `lookupIdent(literal)` to get the correct token type (e.g., `token.LET` or `token.IDENT`).
    - Return the full token.
  - If `isDigit(l.ch)` is true:
    - Call `literal := l.readNumber()`.
    - Return a `token.INT` with the literal.
  - If `l.ch` is a single-character symbol (`+`, `(`, etc.):
    - Directly create and return the token.
  - If it's an unrecognized character:
    - Return an `ILLEGAL` token.

- **Recursive Flaw:** I initially tried to recursively call `getToken` or `readChar` to handle the tokenization, which was incorrect. The correct pattern is to have a single entry point (`NextToken`) and helper functions that **consume and return** the full literal, with `NextToken` then creating the final token object.

- **Upcoming Steps:** The plan is to fully implement the `readIdentifier()`, `readNumber()`, and `lookupIdent()` methods. This will complete the core lexer functionality and allow me to move on to the parser.

## 8th of September 2025
### Lexer Development

  - **Handling Multi-Character Tokens:** The primary focus today was on making the lexer correctly handle tokens that are more than one character long, specifically identifiers and numbers. My previous NextToken function only handled single-character operators, but now it needs to group characters like let and five into a single token.

  - **`readIdentifier()` Function:**
      - Realized that to handle multi-character tokens, I need a helper function that reads a sequence of characters until a non-letter or non-number is found.
      - My initial attempts at this function were incorrect due to an off-by-one error, where the loop was always one character behind the lexer's position.
      - After some discussion, I understood that the loop's condition should be based on the lexer's current state (l.ch), and the loop's body should handle advancing the position (l.readChar()). This is the correct, more robust approach.
      - The most efficient way to capture the identifier is to record the starting index and then slice the input string once the loop terminates. This avoids the overhead of appending characters to a slice.

    **Refining NextToken():**
      - The NextToken() function now uses an if/else if chain to check the type of the current character.
      - If l.ch is a letter, it calls `readIdentifier()`.
      - If l.ch is a number, it will call a similar `readNumber()` function (yet to be implemented).
      - My implementation had an issue where it was returning an empty token for five, which I correctly diagnosed as a missing else case to handle generic identifiers.

    **Handling Keywords:**
      - To differentiate between a keyword (let) and a regular identifier (five), I need a keyword lookup table.
      - I created a var keywords map in the token package, which is a great place to store all language keywords and their corresponding TokenType.
      - This allows for a clean `lookupIdent()` helper function that checks the map and returns either a specific keyword type (token.LET) or a generic identifier type (token.IDENT).

    **Skipping Whitespace:**
      - My tests were failing because the lexer was trying to process spaces.
      - This highlighted the need for a `skipWhitespace()` helper function. This function will contain a simple loop to advance the lexer's position past any space, tab, or newline characters.
      - Calling `skipWhitespace()` at the beginning of `NextToken()` will ensure the lexer is always positioned at the start of a meaningful token. This is a crucial step for a robust lexer.

---

## 9th of September 2025

### Debugging the Lexer

  **Handling Whitespace:**
  - Figured out how to correctly handle all types of whitespace, including tabs and spaces. This fixed the initial error where the lexer was expecting a `LET` token but found a blank one.

  **Preventing Double-Reads:**
  - Identified the off-by-one bug where the lexer would read certain characters twice. The fix was to make the `readChar` function responsible for advancing the state, ensuring that `NextToken` always starts with a fresh character.

  **Recognizing Keywords:**
  - Successfully implemented the logic to check if a valid identifier, like `fn`, is a reserved keyword. This fixed the last remaining issue where the lexer was incorrectly returning an `IDENT` token instead of a `FUNCTION` token.

 - Test for input:
    `let five = 5;`
    `let ten = 10;`
    `let add = fn(x, y) {`
    `x + y;`
    `};`
    `let result = add(five, ten);`
  PASSED
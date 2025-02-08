# monkE Interpreter

This project is an implementation of an interpreter for the `monkE` programming language, inspired by the book **"Writing an Interpreter in Go"** by Thorsten Ball. 

## Overview

`monkE` is a simple, dynamically-typed programming language. This project is a **work in progress** and aims to implement the core features of the language as described in the book. The interpreter is written in Go.

## Features

- **Lexer**: Tokenizes the input source code.
- **Parser**: Parses the tokens into an Abstract Syntax Tree (AST).
- **REPL**: A Read-Eval-Print Loop for interactive programming.

## Project Structure

- `ast/`: Contains the Abstract Syntax Tree (AST) definitions.
- `lexer/`: Contains the lexer implementation.
- `parser/`: Contains the parser implementation.
- `repl/`: Contains the REPL implementation.
- `token/`: Contains token definitions.
- `main.go`: Entry point of the interpreter.

## Getting Started

### Prerequisites

- Go 1.23.5 or higher

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/monkE.git
    cd monkE
    ```

2. Build the project:
    ```sh
    go build
    ```

3. Run the REPL:
    ```sh
    ./monkE
    ```

## Usage

Once the REPL is running, you can start typing `monkE` code and see the results immediately. For example:

```sh
This is monkE programming language!
Feel free to type in Commands
>> 5 * 5 + 10
35
>> 3 + 4 * 5 == 3 * 1 + 4 * 5
true
>> 5 * 10 > 40 + 9
true
>> (10 + 2) * 30 == 300 + 20 * 3
true
>> (5 > 5 == true) != false                                                                                                                                                                  
false
>> 500 / 2 != 250
false
>> if (5 * 5 + 10 > 34) { 99 } else { 100 }
99
>> if ((1000 / 2) + 250 * 2 == 1000) { 9999 }
9999
>> let a = 5;
>> let b = a > 3;
>> let c = a * 99;
>> if (b) { 10 } else { 1 };
10
>> let d = if (c > a) { 99 } else { 100 };
>> d
99
>> d * c * a;
245025
>> let addTwo = fn(x) { x + 2; };
>> addTwo(2)
4
>> let multiply = fn(x, y) { x * y; };
>> multiply(50 / 2, 1 * 2);
50
>> fn(x) { x == 10 }(5)
false
>> fn(x) { x == 10 }(10) 
true
>> let newAdder = fn(x) { fn(y) { x + y } };
>> let addTwo = newAdder(2);
>> addTwo(3);
5
>> let addThree = newAdder(3);
>> addThree(10);
13
>>
```

## License

This project is licensed under the MIT License.

## Acknowledgements

- Thorsten Ball for his excellent book **"Writing an Interpreter in Go"**.

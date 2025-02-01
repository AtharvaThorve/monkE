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
>> let x = 5;
>> x + 10;
15
```

## License

This project is licensed under the MIT License.

## Acknowledgements

- Thorsten Ball for his excellent book **"Writing an Interpreter in Go"**.

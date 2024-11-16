# An Interpreter Written in Go

It is a simple interpreter written in Go, named **Monkey**.  
It was built as a way to learn how interpreted languages are implemented.

## Overview

Monkey is a dynamically typed, interpreted language with a C-like syntax. It supports variable bindings, integers, booleans, arithmetic expressions, built-in functions, first-class and higher-order functions, closures, strings, arrays, and hashes.

### Language Syntax

#### Variables
```monkey
let age = 25;
let name = "John";
```

#### Arithmetic Expressions
```monkey
let sum = 5 + 5;
let product = 5 * 5;
let quotient = 10 / 2;
let difference = 10 - 5;
```

#### Boolean Expressions
```monkey
let isAdult = age > 18;
let isEqual = (10 == 10);
let isNotEqual = (10 != 9);
```

#### Functions
```monkey
let add = fn(x, y) { x + y };
let result = add(5, 10);
```

#### Strings
```monkey
let greeting = "Hello, " + name + "!";
```

#### Arrays
```monkey
let numbers = [1, 2, 3, 4, 5];
let firstNumber = numbers[0];
```

#### Hashes
```monkey
let person = {"name": "John", "age": 25};
let personName = person["name"];
```

### Example Programs

#### Hello World
```monkey
puts("Hello, World!");
```

#### Map Function
```monkey
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            accumulated
        } else {
            iter(rest(arr), push(accumulated, f(first(arr))));
        }
    };

    iter(arr, []);
};

let double = fn(x) { x * 2 };
let result = map([1, 2, 3, 4], double);
puts(result); // [2, 4, 6, 8]
```

#### Factorial
```monkey
let factorial = fn(n) {
    if (n == 0) {
        1
    } else {
        n * factorial(n - 1);
    }
};

puts(factorial(5)); // 120
```

#### Fibonacci
```monkey
let fibonacci = fn(n) {
    if (n == 0) {
        0
    } else if (n == 1) {
        1
    } else {
        fibonacci(n - 1) + fibonacci(n - 2);
    }
};

puts(fibonacci(10)); // 55
```

### Built-in Functions

#### `len`
Returns the length of a string or array.
```monkey
len("hello"); // 5
len([1, 2, 3]); // 3
```

#### `first`
Returns the first element of an array.
```monkey
first([1, 2, 3]); // 1
```

#### `last`
Returns the last element of an array.
```monkey
last([1, 2, 3]); // 3
```

#### `rest`
Returns the array without the first element.
```monkey
rest([1, 2, 3]); // [2, 3]
```

#### `push`
Adds an element to the end of an array.
```monkey
push([1, 2, 3], 4); // [1, 2, 3, 4]
```

## How to Use
To run it, you'll need [Git](https://git-scm.com) and [Go](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/ekastn/monkey

# Go into the repository
$ cd monkey

# Build the project
$ go build -o bin/monkey cmd/main.go

# Run the REPL
$ ./bin/monkey
```

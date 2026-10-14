# Hello World

## What this topic is
The traditional first program in any programming language that prints "Hello, World!" to the console. This demonstrates the basic structure of a Go program.

## When to use it
- Learning the basics of Go syntax
- Verifying your Go installation is working correctly
- Starting any new Go project

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Common mistakes
- Forgetting the `package main` declaration
- Forgetting to import the `fmt` package for printing
- Using `Println` without the `fmt` prefix
- Missing semicolons (Go automatically inserts them, but understanding when they're needed is important)

## Command to run it
```bash
go run main.go
```
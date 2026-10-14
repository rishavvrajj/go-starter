# Packages

## What this topic is
Packages in Go are used to organize and reuse code. Every Go file belongs to a package, and packages help with encapsulation, dependency management, and building modular applications. The `main` package is special - it defines an executable program.

## When to use it
- Organizing related functionality into cohesive units
- Creating reusable libraries that can be shared across projects
- Encapsulating implementation details (hiding complexity)
- Managing dependencies and avoiding naming conflicts
- Building large applications with separated concerns
- Creating public APIs that others can use

## Minimal syntax example
```go
// File: mathutil/add.go
package mathutil

// Exported function (starts with capital letter)
func Add(a, b int) int {
    return a + b
}

// Unexported function (starts with lowercase letter)
func subtract(a, b int) int {
    return a - b
}

// File: main.go
package main

import (
    "fmt"
    "example.com/mathutil" // local or remote package
)

func main() {
    result := mathutil.Add(5, 3) // 8
    fmt.Println(result)
    
    // mathutil.subtract(10, 4) // ❌ compile error: not exported
}
```

## Common mistakes
- Forgetting that package names should match the directory name
- Not understanding exported vs unexported identifiers (capitalization controls visibility)
- Having multiple packages in the same directory (all files must share the same package name)
- Using relative paths in imports (should use module paths)
- Forgetting to run `go mod tidy` after adding/removing imports
- Circular dependencies (package A imports B, B imports A)
- Not initializing packages properly with init() functions
- Forgetting that package initialization happens in dependency order
- Trying to access unexported fields or functions from another package
- Using blank identifiers incorrectly when importing (only for side-effects)

## Command to run it
```bash
go run main.go
```
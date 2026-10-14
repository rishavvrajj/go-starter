# Go Learning & Quick Reference Repository

Welcome to your comprehensive Go (Golang) learning repository and personal reference! This repository has been structured in an ordered, beginner-friendly learning path, designed for quick lookups when you need a fast reminder of Go syntax, collections, structs, interfaces, generics, or concurrency.

---

## 🚀 Recommended Learning Path

| # | Topic | Description | Key Concepts |
|---|-------|-------------|--------------|
| 01 | [Hello World](01-hello-world/) | Introduction to Go programs | `package main`, `import`, `fmt.Println` |
| 02 | [Values](02-values/) | Basic value types | Integers, floats, booleans, strings, runes |
| 03 | [Variables](03-variables/) | Variable declarations & scope | `var`, `:=`, zero values, block scope |
| 04 | [Constants](04-constants/) | Immutable values | `const`, typed vs untyped constants, iota |
| 05 | [If/Else](05-if-else/) | Conditional branching | `if`, `else`, initialization statements |
| 06 | [Switch](06-switch/) | Multi-way branching | `switch`, `case`, type switches, `fallthrough` |
| 07 | [For Loops](07-for-loops/) | Repetition and iteration | `for`, condition-only loops, infinite loops |
| 08 | [Range](08-range/) | Iterating over collections | `range` over slices, maps, strings, channels |
| 09 | [Functions](09-functions/) | First-class functions | Multiple returns, function types, closures |
| 10 | [Closures](10-closures/) | Function state capture | Capturing variables, counter pattern, memoization |
| 11 | [Variadic Functions](11-variadic-functions/) | Variable argument lists | `...T` parameters, expanding slices |
| 12 | [Arrays](12-arrays/) | Fixed-size sequences | Fixed length, value semantics, index access |
| 13 | [Slices](13-slices/) | Dynamic array views | `make`, `append`, `len`, `cap`, reference semantics |
| 14 | [Maps](14-maps/) | Key-value associative structures | `make`, comma-ok idiom, `delete`, iteration |
| 15 | [Pointers](15-pointers/) | Memory addresses | `&` (address-of), `*` (dereference), passing by pointer |
| 16 | [Structs](16-structs/) | Composite data types | Struct literals, methods, value/pointer receivers, embedding |
| 17 | [Interfaces](17-interfaces/) | Behavioral contracts | Implicit implementation, polymorphism, dependency injection |
| 18 | [Enums](18-enums/) | Type-safe enumerations | Custom types, `iota`, `Stringer` interface |
| 19 | [Generics](19-generics/) | Type-parameterized code | Type parameters `[T any]`, constraints, comparable |
| 20 | [Packages](20-packages/) | Modular code organization | Exported vs unexported identifiers, `init()`, modules |
| 21 | [Goroutines](21-goroutines/) | Lightweight concurrency | `go` keyword, `sync.WaitGroup` |
| 22 | [Channels](22-channels/) | Goroutine communication | Unbuffered/buffered channels, `select`, closing |
| 23 | [Mutex](23-mutex/) | Thread synchronization | `sync.Mutex`, `sync.RWMutex`, `sync.Once` |
| 24 | [Files](24-files/) | File I/O and OS operations | Reading, writing, `bufio`, directories, paths |

---

## 🏃‍♂️ How to Run Any Example

Every numbered folder contains a runnable `main.go` file (or subpackages where appropriate). You can run any example directly using `go run`:

```bash
# Navigate to any topic directory
cd 01-hello-world

# Run the program
go run main.go
```

For the packages example (`20-packages`), run it from its directory:
```bash
cd 20-packages
go run main.go
```

---

## 🛠️ Go Version & Module Setup

- **Go Version:** 1.22+ (configured for Go 1.27.1)
- **Module Name:** `github.com/rishavvrajj/go-starter`

### Module Commands
```bash
# Initialize module (already done)
go mod init github.com/rishavvrajj/go-starter

# Tidy dependencies
go mod tidy

# Build all packages in the module
go build ./...

# Run static analysis (vet)
go vet ./...

# Run tests (if any)
go test ./...
```

---

## 💡 Quick Reference Tips
- **Exported identifiers** start with an uppercase letter (`Name`), making them public.
- **Unexported identifiers** start with a lowercase letter (`name`), keeping them package-private.
- **Short variable declaration (`:=`)** can only be used inside functions.
- **Slices and maps** are reference types; **arrays and structs** are value types.
- Always use **`defer file.Close()`** after opening files and **`defer mu.Unlock()`** after locking mutexes.

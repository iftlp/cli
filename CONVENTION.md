# Go Conventions

## 1. Naming Conventions

### 1.1 Package Naming

- Use **lowercase** and short, meaningful names.
- Avoid underscores (`_`) or mixed caps.

```go
package customer
```

### 1.2 Struct Naming

- Use **PascalCase**.
- Name structs explicitly and meaningfully.

```go
type CustomerManager struct {
    // Fields
}
```

### 1.3 Function Naming

- Use **PascalCase** for exported functions.
- Use **camelCase** for unexported functions.

```go
func CalculateTotal() int {
    return 0
}

func processOrder() {
    // Internal function
}
```

### 1.4 Variable and Field Naming

- Use **camelCase** for local variables.
- Use **PascalCase** for exported struct fields.
- Use **short, clear, and meaningful** names.

```go
var customerName string

// Inside a struct

type Customer struct {
    Name  string  // Exported field
    email string  // Unexported field
}
```

### 1.5 Constant Naming

- Use **PascalCase** for exported constants.
- Use **camelCase** for unexported constants.

```go
const MaxRetryCount = 5
const apiKey = "12345"
```

### 1.6 Interface Naming

- Use `-er` suffix when possible.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type CustomerRepository interface {
    Save(customer Customer) error
}
```

### 1.7 Error Naming

- Define errors using `errors.New`.
- Use `Err` prefix.

```go
var ErrNotFound = errors.New("not found")
```

## 2. Comments and Documentation

### 2.1 Using GoDoc Comments

- Use `//` for single-line comments.
- Use full sentences starting with the function name.

```go
// CustomerManager handles customer operations.
type CustomerManager struct {}

// AddCustomer adds a new customer.
func (cm *CustomerManager) AddCustomer(c Customer) {
    // Implementation
}
```

## 3. Unit Tests

### 3.1 Using Testing Package

- Use Go's built-in `testing` package.
- Name test functions as `Test[Function]`.
- Use `*_test.go` suffix for test files.

```go
package customer_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAddCustomer(t *testing.T) {
    cm := customer.CustomerManager{}
    c := customer.Customer{Name: "John Doe"}
    
    err := cm.AddCustomer(c)
    
    assert.NoError(t, err)
}
```

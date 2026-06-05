# Go Cheat Sheet

A quick reference for Go syntax, aimed at developers coming from Python.

---

## Variables

```go
// declare and assign (type is inferred)
x := 42
name := "hello"

// explicit type
var x int = 42

// multiple assignment
a, b := 1, 2
```

---

## Types

```go
bool
int, int8, int16, int32, int64
float32, float64
string
```

---

## Constants

```go
const Pi = 3.14

// iota auto-increments in a const block
const (
    Empty = iota  // 0
    Black         // 1
    White         // 2
)
```

---

## Functions

```go
// basic
func add(a int, b int) int {
    return a + b
}

// multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// calling a function with multiple returns
result, err := divide(10, 2)
if err != nil {
    // handle error
}
```

---

## Structs

```go
// define
type Point struct {
    Row int
    Col int
}

// create
p := Point{Row: 3, Col: 4}
p := Point{3, 4}  // shorthand if order matches

// access fields
p.Row
p.Col
```

---

## Methods

```go
// method on a struct (value receiver — does not modify original)
func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.Row, p.Col)
}

// method on a pointer (can modify the original)
func (b *Board) Set(p Point, c Color) {
    b.Cells[p.Row][p.Col] = c
}

// calling a method
b.Set(p, Black)
```

**Rule of thumb:** use a pointer receiver (`*Type`) when the method needs to modify the struct, or when the struct is large.

---

## Pointers

```go
x := 42
p := &x    // p is a pointer to x
*p = 99    // dereference to set the value — x is now 99

// structs are often passed as pointers
func doSomething(b *Board) {
    b.Size  // Go auto-dereferences, no need for b->Size like C
}
```

---

## Slices

```go
// create
nums := []int{1, 2, 3}

// make (empty slice with length)
cells := make([]int, 9)

// append
nums = append(nums, 4)

// length
len(nums)

// index
nums[0]

// slice of slice
nums[1:3]  // elements at index 1 and 2
```

---

## 2D Slices

```go
// create a 9x9 grid
grid := make([][]int, 9)
for i := range grid {
    grid[i] = make([]int, 9)
}

// access
grid[row][col]
```

---

## Maps

```go
// create
scores := map[string]int{
    "Black": 0,
    "White": 0,
}

// set
scores["Black"] = 10

// get
s := scores["Black"]

// make (empty map)
scores := make(map[string]int)
```

---

## Control Flow

```go
// if/else (no parentheses)
if x > 0 {
    // ...
} else if x < 0 {
    // ...
} else {
    // ...
}

// for (the only loop in Go)
for i := 0; i < 10; i++ { }

// while-style
for x < 100 { }

// range over slice
for i, val := range nums { }
for _, val := range nums { }  // ignore index

// range over map
for key, val := range scores { }
```

---

## Error Handling

```go
// functions signal errors via a second return value
func (g *Game) PlaceStone(p Point) error {
    if !InBounds(g.Board, p) {
        return errors.New("point is out of bounds")
    }
    return nil  // nil means no error
}

// always check errors at the call site
err := g.PlaceStone(p)
if err != nil {
    fmt.Println("invalid move:", err)
}
```

---

## Zero Values

Every type has a default zero value — no null pointer surprises like Python's `None`.

| Type | Zero value |
|---|---|
| `int`, `float64` | `0` |
| `bool` | `false` |
| `string` | `""` |
| pointer, slice, map | `nil` |

---

## Packages & Imports

```go
// top of every file
package game

// importing
import (
    "fmt"
    "errors"
)

// using
fmt.Println("hello")
errors.New("something went wrong")
```

---

## Useful fmt Verbs

```go
fmt.Println("hello")          // print with newline
fmt.Printf("val: %d\n", 42)  // formatted print

// common verbs
%d   // integer
%f   // float
%s   // string
%v   // any value (default format)
%T   // type of value
```

---

## Testing

```go
func TestSomething(t *testing.T) {
    result := MyFunction(5)
    if result != 10 {
        t.Errorf("expected 10, got %d", result)
    }
}
```

- `t.Errorf` — marks test failed, continues running
- `t.Fatalf` — marks test failed, stops immediately

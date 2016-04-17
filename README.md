# GoCount

This is a simple package with just one purpose: count the number of running
goroutines that are not runtime goroutines. It is different from
[`runtime.NumGoroutine`](https://godoc.org/runtime#NumGoroutine), as that
function counts all goroutines (including e.g. garbage collection goroutines).

Example:

```go
package main

import (
    "github.com/aykevl/gocount"
    "fmt"
)

func main() {
    fmt.Println("Number of running goroutines:", gocount.Number())

    block := make(chan struct{})
    done := make(chan struct{})

    fmt.Print("hello, ")

    go func() {
        fmt.Println("world")
        done <- struct{}{}
        <-block
    }()
    <-done

    fmt.Println("Number of running goroutines:", gocount.Number())
}
```

Output:

    Number of running goroutines: 1
    hello, world
    Number of running goroutines: 2


The package is licensed under the 3-clause BSD license.

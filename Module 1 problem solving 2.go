package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Scan(&a)
    b = a % 10 * 100 + a % 100 / 10 * 10 + a / 100
    fmt.Println(b)
}

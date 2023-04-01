package main

import "fmt"

func main() {
    var n int
    k := 1
    fmt.Scan(&n)
    fmt.Print(k)
    fmt.Print(" ")
    for k < n {
        k = k * 2
        if (k <= n) {
            fmt.Print(k)
            fmt.Print(" ")
        }
    }
}
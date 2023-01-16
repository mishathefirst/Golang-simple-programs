package main

import (
    "fmt"
)

func main () {
    var a, b int
    var c float64
    fmt.Scan(&a, &b)
    c = (float64(a) + float64(b)) / 2
    fmt.Println(c)
}


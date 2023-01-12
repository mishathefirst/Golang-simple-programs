package main

import (
    "fmt"
)

func main() {
    var a, sum int
    fmt.Scan(&a)
    sum = a % 10 + a % 100 / 10 + a / 100
    fmt.Println(sum)
}

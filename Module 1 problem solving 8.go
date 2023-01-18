package main

import (
    "fmt"
)

func main() {
    var n, curr, sum int
    min := 10000000000
    fmt.Scan(&n)
    for i:=0; i < n; i++ {
        fmt.Scan(&curr)
        if curr < min {
            min = curr
            sum = 1
        } else if curr == min {
            sum++
        }
    }
    fmt.Println(sum)
}


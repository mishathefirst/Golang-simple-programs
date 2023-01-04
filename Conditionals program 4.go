package main

import (
    "fmt"
)

func main() {
    var a, sum1, sum2 int
    fmt.Scan(&a)
    
    sum1 = a / 100000 + a / 10000 % 10 + a / 1000 % 10
    sum2 = a % 10 + a % 100 / 10 + a % 1000 / 100
    
    if sum1 == sum2 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

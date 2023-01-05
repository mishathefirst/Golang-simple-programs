package main

import "fmt"

func main() {
    var a, b, sum int
    fmt.Scan(&a, &b)
    for i:=a; i < b + 1; i++ {
        sum = sum + i
    }
    fmt.Println(sum)
}


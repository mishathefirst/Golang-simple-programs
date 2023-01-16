package main

import (
    "fmt"
)

func main() {
    var sum, n, curr int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&curr)
        if curr == 0 {
            sum++ 
        }
    }
    fmt.Println(sum)
}




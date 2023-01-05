package main

import (
    "fmt"
)

func main() {
    var (
        number int = 1
        max int
        sum int
    )
    for number != 0 {
        fmt.Scan(&number)
        if number > max {
            max = number
            sum = 0
        }
        if number == max {
            sum++
        }
    }
    fmt.Println(sum)
}



package main

import "fmt"

func main() {
    var x, p, y, years int
    
    fmt.Scan(&x, &p, &y)
    p = p + 100
    x = x * 100
    y = y * 100
    for x < y {
        x = x * p / 100
        years++
    }
    fmt.Println(years)
}


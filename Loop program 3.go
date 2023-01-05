package main

import "fmt"

func main() {
    var amount, number, sum int
    fmt.Scan(&amount)
    for i:=0; i < amount; i++ {
        fmt.Scan(&number)
        if number >= 10 && number < 100 && number % 8 == 0 {
            sum = sum + number
        }
    }
    fmt.Println(sum)
}



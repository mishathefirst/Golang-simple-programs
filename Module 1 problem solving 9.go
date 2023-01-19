package main

import (
    "fmt";
    "math"
)

func main() {
    var a, sum int
    fmt.Scan(&a)
    sum = a % 10
    for i:=2; i < 8; i++ {
        if a / int(math.Pow(float64(10), float64(i))) != 0 {
            sum += a % int(math.Pow(float64(10), float64(i))) / int(math.Pow(float64(10), float64(i - 1)))
        } else {
            sum += a % int(math.Pow(float64(10), float64(i))) / int(math.Pow(float64(10), float64(i - 1)))
            break
        }
    }
    sum = sum % 10 + sum / 10
    fmt.Println(sum)
}



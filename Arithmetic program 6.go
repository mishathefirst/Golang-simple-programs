package main

import "fmt"

func main() {
    var degrees, hours, minutes int
    fmt.Scan(&degrees)
    hours = degrees / 30
    minutes = 2 * (degrees % 30)
    fmt.Print("It is ", hours, " hours ", minutes, " minutes.")
}



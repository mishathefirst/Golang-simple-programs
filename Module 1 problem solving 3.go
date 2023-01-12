package main

import (
    "fmt"
)

func main() {
    var sum, hours, minutes int
    fmt.Scan(&sum)
    hours = sum / 3600
    minutes = (sum - hours * 3600) / 60
    fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}

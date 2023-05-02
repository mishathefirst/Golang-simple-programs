package main

import (
    "fmt"
    "time"
)

func main() {
    var line string
    fmt.Scan(&line)
    convertingTime, err := time.Parse(time.RFC3339, line)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(convertingTime.Format(time.UnixDate))
}
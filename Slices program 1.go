package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    a := make([]int, n, n)
    for i:=0; i < n; i++ {
        fmt.Scan(&a[i])
    }
    fmt.Println(a[3])
}
func minimumFromFour() int {
    var min int
    var a [4]int
    for i:=0; i < 4; i++ {
        fmt.Scan(&a[i])
    }
    min = a[0]
    for i:=1; i < 4; i++ {
        if a[i] < min {
            min = a[i]
        }
    }
    return min
}


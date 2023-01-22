func sumInt(a ...int) (int, int) {
    length := len(a)
    var sum int
    for i:=0; i < len(a); i++ {
        sum += a[i]
    }
    return length, sum
}

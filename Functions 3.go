func fibonacci(n int) int {
    var prevres, res, a int
    if n == 1 || n == 2 {
        return 1
    } else {
        a = 1
        prevres = 1
        res = 1
        for i:=0; i < n - 2; i++ {
            prevres = res
            res += a
            a = prevres
        }
    }
    return res
}


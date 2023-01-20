func vote(x int, y int, z int) int {
    var zero, one int
    if x == 0 {
        zero++
    } else {
        one++
    }
    if y == 0 {
        zero++
    } else {
        one++
    }
    if z == 0 {
        zero++
    } else {
        one++
    }
    if zero > one {
        return 0
    } else {
        return 1
    }
}




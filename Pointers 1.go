func test(x1 *int, x2 *int) {
    temp := *x1
    *x1 = *x2
    *x2 = temp
    //x1 = &(*x2)
    //x2 = &(*x1)
    fmt.Println(*x1, *x2)
}


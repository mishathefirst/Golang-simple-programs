var a int
newmap := make(map[int]int)
for i:=0; i < 10; i++ {
    fmt.Scan(&a)
    if _, ok := newmap[a]; ok {
        fmt.Print(newmap[a], " ") 
    } else {
        newmap[a] = work(a)
        fmt.Print(newmap[a], " ")
    }
}


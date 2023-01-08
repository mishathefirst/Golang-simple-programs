var workArray [10]uint8
var first, second int
//var temp byte
for i:=0; i < 10; i++ {
    fmt.Scan(&workArray[i])
}
for i:= 0; i < 3; i++ {
    fmt.Scan(&first)
    fmt.Scan(&second)
    //temp = workArray[first]
    //workArray[first] = workArray[second]
    //workArray[second] = temp
    workArray[first], workArray[second] = workArray[second], workArray[first]
}
for i:=0; i < 10; i++ {
    fmt.Print(workArray[i], " ")
}


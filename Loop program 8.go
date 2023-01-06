package main

import (
    "fmt";
    "math"
)

func main() {
    var a, b, curr, coeff, numb1, numb2, curr2, coeff2 int
    fmt.Scan(&a, &b)
    coeff = 1
    for i:=1; i <= 4; i++ {
        coeff = coeff * 10
        numb1++
        if a / coeff == 0 {
            break
        }
    }
    coeff = 1
    for i:=1; i <= 4; i++ {
        coeff = coeff * 10
        numb2++
        if b / coeff == 0 {
            break
        }
    }
    for i:=0; i < numb1; i++ {
        coeff = int(math.Pow10(numb1 - (i + 1)))
        curr = a / coeff % 10
        for j:=0; j < numb2; j++ {
            coeff2 = int(math.Pow10(numb2 - (j + 1)))
            curr2 = b / coeff2 % 10
            if curr == curr2 {
                fmt.Print(curr, " ")
                break
            }
        }
    }
}
    


package main

import "fmt"

func CheckEvenOrOdd(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
}

func main() {
    n := 8

    result := CheckEvenOrOdd(n)
    fmt.Println(result)
}

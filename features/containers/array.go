package main

import (
    "fmt"
)

func test0() {
    values := make([]int, 0, 0)
    var v []int
    v = append(values, 1)
    v = append(values, 2)
    v = append(values, 3)

    fmt.Println(values)
    fmt.Println(v)
}

func main() {
    test0()
}

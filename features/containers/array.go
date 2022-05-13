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

func UpdateArray(a []int) {
    a = append(a, 10)
}

func test1() {
    arr := make([]int, 0)
    arr = append(arr, 1)
    arr = append(arr, 2)

    fmt.Println(arr)
    UpdateArray(arr)
    fmt.Println(arr)
}

func main() {
    test0()
    test1()
}

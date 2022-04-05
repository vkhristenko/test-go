package main

import (
    "fmt"
)

func test0() {
    m := make(map[int32]int32)
    m[1] = 1
    m[2] = 2
    m[3] = 3

    fmt.Println(m)
    for key, _ := range m {
        if key == 1 {
            delete(m, 1)
        }
    }
    fmt.Println(m)
}

func main() {
    test0()
}

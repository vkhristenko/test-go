package main

import (
    "fmt"
)

func some_func() *int {
    local_var := 10
    return &local_var
}

func test0() {
    v := some_func() 
    fmt.Printf("value = %d\n", *v)
}

func main() {
    test0()
}

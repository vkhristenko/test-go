package main

import (
    "fmt"
)

type Base struct {
    b int
}

type Container struct {
    Base
    c string
}

func test0() {
    c := Container {
        Base: Base {
            b: 10,
        },
        c: "hello world",
    }
    fmt.Println(c)
    fmt.Printf("b = %d c = %s\n", c.b, c.c)
    fmt.Printf("b = %d c = %s\n", c.Base.b, c.c)
}

func main() {
    test0()
}

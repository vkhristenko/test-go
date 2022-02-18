package main

import (
    "fmt"
)

type Data struct {
    data string
}

func Test() string {
    s := "hello"
    fmt.Printf("address = %p\n", &s)
    return s
}

func main() {
    fmt.Println("hello from bin1")
    SomeFunc()
    s := Test()
    fmt.Printf("address = %p\n", &s)
}

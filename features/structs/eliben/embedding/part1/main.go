package main

import (
    "fmt"
    "sync"
)

type Base struct {
    b int
}

type Container struct {
    Base
    c string
}

func (this Base) Describe() string {
    return fmt.Sprintf("base %d belongs to us", this.b)
}

func (this Container) Describe() string {
    return fmt.Sprintf("co -> {b: %v, c: %v}\n", this.b, this.c)
}

func test0() {
    {
        co := Container {}
        co.b = 1
        co.c = "string"
        fmt.Printf("co -> {b: %v, c: %v}\n", co.b, co.c)
        fmt.Println(co.Base.Describe())
        fmt.Println(co.Describe())
    }

    {
        co := Container{
            Base: Base{
                b:10,
            },
            c: "foo",
        }
        fmt.Printf("co -> {b: %v, c: %v}\n", co.b, co.c)
        fmt.Println(co.Base.Describe())
        fmt.Println(co.Describe())
    }
}

type SessionCache struct {
    sync.Mutex
    m map[string]int32
}

func test1() {
    cache := SessionCache{
        m: map[string]int32{
            "1": 1,
            "2": 2,
            "3": 3,
        },
    }
    cache.Lock()
    cache.Unlock()
}

func main() {
    test0()
    test1()
}

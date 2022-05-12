package main


import (
    "fmt"
)

type FooI interface {
    Foo() string 
}

type Container struct {
    FooI
}

func (this Container) Foo() string {
    return this.FooI.Foo()
}

func sink(f FooI) {
    fmt.Println("sink: ", f.Foo())
}

type FooImpl struct {}
func (this FooImpl) Foo() string {
    return "The FooImpl Foo"
}

func test0() {
    co := Container{
        FooI: FooImpl{},
    }
    sink(co)
}

func main() {
    test0()
}

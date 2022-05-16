package main

import (
    "fmt"
)

type Animal interface {
    Speak() string
}

type Dog struct {}
type Cat struct {}

func (this Dog) Speak() string {
    return "Dog"
}

func (this Cat) Speak() string {
    return "Cat"
}

func test0() {
    animals := []Animal{Dog{}, Cat{}, Dog{}}
    for _, a := range animals {
        fmt.Println(a.Speak())
        fmt.Println(a)
    }
}

func main() {
    test0()
}

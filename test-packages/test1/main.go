package main

import (
	"fmt"
	"github.com/vkhristenko/test-go/test-packages/p1"
	"github.com/vkhristenko/test-go/test-packages/p2"
    "test-packages/p3"
)

func main() {
    fmt.Println(p1.SomeFunc())
    fmt.Println(p2.SomeFunc())
    fmt.Println(p3.SomeFunc())
}

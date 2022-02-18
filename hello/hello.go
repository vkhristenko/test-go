package main

import (
    "fmt"
    "github.com/vkhristenko/test-go/greetings"
)

func main() {
    message := greetings.Hello("Viktor")
    fmt.Println(message)
}

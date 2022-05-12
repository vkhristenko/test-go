package main

import (
    "fmt"
)

type ReaderI interface {
    Read(p []byte) (int, error)
}

type WriterI interface {
    Write(p []byte) (int, error)
}

type ReadWriterI interface {
    ReaderI
    WriterI
}

func test0() {
    fmt.Println("test0")
}

func main() {
    test0()
}

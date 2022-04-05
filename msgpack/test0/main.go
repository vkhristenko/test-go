package main

import (
	"fmt"

	msgpack "github.com/vmihailenco/msgpack/v5"
)

type Data struct {
    X, Y, Z int
    Msg string
    Value1, Value2 float64
}

func main() {
    bytes, err := msgpack.Marshal(&Data {
        X: 1,
        Y: 1,
        Z: 1,
        Msg: "value",
        Value1: 1.1,
        Value2: 2.2,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(len(bytes))
    for i, b := range bytes {
        if i % 20 == 0 {
            fmt.Println()
        }
        fmt.Printf("%02x ", b)
    }
    fmt.Println()

    var data Data
    err = msgpack.Unmarshal(bytes, &data)
    if err != nil {
        panic(err)
    }
    fmt.Println(data)
}

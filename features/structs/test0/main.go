package main

import (
    "fmt"
)

type DataItem struct {
    str string
}

type Data struct {
    arr []DataItem
}

func test0() {
    data := Data{arr: make([]DataItem, 0)}
    fmt.Println(data)
    data.arr = append(data.arr, DataItem{str: "hello"})
    fmt.Println(data)
}

type Data1 struct {
    arr map[string]string
}

func test1() {
    m := make(map[string]Data1)
    m["hello"] = Data1 {
        arr: make(map[string]string),
    }
    fmt.Println(m)
    m["hello"].arr["1"] = "1"
    fmt.Println(m)
}

func main() {
    test0()
    test1()
}

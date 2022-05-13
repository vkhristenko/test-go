package main

import (
    "fmt"
)

func printValueAddress(value int) {
    fmt.Printf("value = %d address = %p\n", value, &value)
}

func test0() {
    fmt.Println("running test 0")

    value := 10
    value1 := value
    printValueAddress(value)
    printValueAddress(value1)
}

type Point struct {
    x int
    y int
}

func (self Point) Shift(value int) {
    self.x += value
    self.y += value
}

func (self *Point) MutShift(value int) {
    self.x += value
    self.y += value
}

func printPointValue(p Point) {
    fmt.Println("point = ", p)
    fmt.Printf("address = %p\n", &p)
}

func printPointPointer(p *Point) {
    fmt.Println("point = ", *p)
    fmt.Printf("address = %p\n", p)
}

func shiftPoint(p *Point, shift int) {
    p.x += shift
    p.y += shift
}

func test1() {
    p := Point {1, 2}
    printPointValue(p)
    p.x = 10
    printPointValue(p)

    printPointPointer(&p)
    printPointPointer(&p)
    shiftPoint(&p, 10)
    printPointPointer(&p)
    p.Shift(10)
    printPointPointer(&p)
    p.MutShift(10)
    printPointPointer(&p)
}

func updateMap[K int, V int](m map[K]V) {
    m[5] = 5
}

func test2() {
    m := map[int]int{
        1: 2,
        2: 2,
        3: 3,
    }
    fmt.Println("map: ", m)
    updateMap(m)
    fmt.Println("map: ", m)
}

type TestCopying struct {
    Ptr *int32
}

func PrintTestCopying(s TestCopying) {
    fmt.Printf("value = %d\n", *s.Ptr)
}

func ModTestCopying(s TestCopying) {
    *s.Ptr += 10
}

type Data1 struct {
    str string
}

func test3() {
    fmt.Println("*** test3 ***")

    s := TestCopying {
        Ptr: new(int32),
    }
    *s.Ptr = 10
    PrintTestCopying(s)
    ModTestCopying(s)
    PrintTestCopying(s)
    ptr := &(*s.Ptr)
    fmt.Printf("addr = %p\n", ptr)
    *ptr += 30
    PrintTestCopying(s)

    data := Data1 {
        str: "some string",
    }
    pstr := &data.str
    fmt.Printf("member field string address = %p\n", pstr)
}

func main() {
    test0()
    test1()
    test2()
    test3()
}

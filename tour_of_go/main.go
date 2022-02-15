package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func add_ints(x int, y int) int {
	return x + y
}

func get_ints(x, y int) (int, int) {
	return x, y
}

func test_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func test_switch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("something else... ", os)
	}
}

func test_if() int {
    x := 10
    if v := x+1; v < 10 {
        return v
    } else {
        return v+10
    }
}

func test_pointer() {
    x := 10
    p := &x
    fmt.Printf("x = %d p = %p\n", x, p)
}

type Point struct {
    x int
    y int
}

func test_point() {
    p := Point{1, 2}
    p.x = 10

    fmt.Printf("x = %d y = %d\n", p.x, p.y)
    fmt.Println("point = ", p)
}

func print_array(s []int) {
    fmt.Printf("slice address = %p\n", &s)
    fmt.Printf("slice length = %d\n", len(s))
    fmt.Printf("slice capacity = %d\n", cap(s))
    fmt.Println(s)
}

func test_array() {
    var a [10]int
    for i:=0; i<10; i++ {
        a[i] = i
    }

    for i:=0; i<10; i++ {
        fmt.Printf("a[ %d ] = %d\n", i, a[i])
    }

    slice := a[1:9]
    for i:=0; i<len(slice); i++ {
        fmt.Println(slice[i])
    }

    slice1 := a[:]
    fmt.Printf("a's address = %p\n", &a)
    fmt.Printf("slice1 address = %p\n", slice1)
    fmt.Printf("newly created slice = %p\n", a[:])
    print_array(slice1)
    print_array(a[:])
    print_array(a[1:])
    print_array(a[1:4][1:])
    var s []int = nil
    fmt.Println(s)

    {
        s := make([]int, 0, 10)
        fmt.Println(s)
        print_array(s[:cap(s)])
        s = append(s, 1, 2, 3, 4)
        print_array(s)
        print_array(s[:cap(s)])

        v := make([]int, 0, 10)
        print_array(v)
        v = append(v, 1,2,3)
        print_array(v)
        for i, v := range v {
            fmt.Println(i, v)
        }
    }
}

func TestMap() {
    var m = map[string]int{
        "1" : 1,
        "2" : 2,
    }
    m["3"] = 3
    fmt.Println(m)
}

type Point3d struct {
    x, y, z float64
}

func (p Point3d) Add(rhs Point3d) Point3d {
    p = Point3d {
        p.x + rhs.x,
        p.y + rhs.y,
        p.z + rhs.z,
    }

    return p
}

func (p *Point3d) AddConstant(x float64) Point3d {
    (*p) = Point3d {
        p.x + x,
        p.y + x,
        p.z + x,
    }

    return (*p) 
}

func Pass(p *Point3d) {
    fmt.Printf("Pass point address = %p\n", p)
}

func TestMethod() {
    p := Point3d {1, 2, 3}
    p = p.Add(Point3d{1,2,3})
    fmt.Println(p)

    fmt.Printf("point address = %p\n", &p)
    Pass(&p)

    p1 := Point3d{1,2,3}
    p2 := p1.AddConstant(10.5)
    fmt.Printf("@p1 = %p\n", &p1)
    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Printf("@p2 = %p\n", &p2)
}

func TestSlice() {
    a := make([]int, 10, 10)
    fmt.Println(a)
    for i := range a {
        a[i] = i
    }
    fmt.Println(a[1:3])
}

func Adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func TestClosures() {
    pos, neg := Adder(), Adder()
    for i:=0; i<10; i++ {
        fmt.Println(pos(i), neg(-2*i))
    }
}

type Something struct {
    x, y uint64
}

func TakePointerToSomething(p *Something) {
    fmt.Printf("something at %p\n", p)
}

func TestSomething() {
    s := Something{}
    p := &s
    TakePointerToSomething(p)
    TakePointerToSomething(&s)
}

func (s *Something) Add(v uint64) {
    fmt.Printf("address of s = %p\n", s)
    s.x += v
    s.y += v
    fmt.Println(s)
}

func AddFunc(s *Something, v uint64) {
    s.x += v
    s.y += v
}

func TestMethodsAgain() {
    s := Something{}
    fmt.Printf("address of s = %p\n", &s)
    fmt.Println(s)
    s.Add(10)
    fmt.Println(s)
    fmt.Printf("address of s = %p\n", &s)

    AddFunc(&s, 10)
    fmt.Println(s)
}

type Abser interface {
    Abs() float64
}

func TestInterfaces() {

}

func TakeEmptyInterface(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("integer: %d\n", v)
    case string:
        fmt.Printf("string %s\n", v)
    default:
        fmt.Printf("default case %T\n", i)
    }
}

func TestTypeSwitch() {
    var i interface{}

    i = 5
    TakeEmptyInterface(i)
    i = "hello"
    TakeEmptyInterface(i)
    i = Something{}
    TakeEmptyInterface(i)
}

type Person struct {
    Name string
    Age int
}

func (p *Person) String() string {
    fmt.Printf("String() Person address is %p\n", p)
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestStringer() {
    p1 := Person{Name: "Viktor", Age: 10}
    fmt.Printf("Person address is %p\n", &p1)
    fmt.Println(&p1)
}

func MyGo(s string) {
    fmt.Println(s)
}

func AddInts(x,y int, ch chan int) {
    value := x+y
    ch <- value
}

func TestGoroutines() {
    go MyGo("some string")
    ch := make(chan int)
    go AddInts(10, 20, ch)
    value := <- ch
    fmt.Println(value)
}

func main() {
	var i int
	var ii int = 5
	var iii = 5
	iiii := 10
	const str = "Some String"

	fmt.Println("Hello world!")
	fmt.Println("value is ", rand.Intn(10))
	fmt.Println(i, ii, iii, iiii)

    //defer fmt.Println("printed last")

	test_loop()
	test_switch()
    test_if()
    test_pointer()
    test_point()
    test_array()
    TestMap()
    TestMethod()
    TestSlice()
    TestClosures()
    TestSomething()
    TestMethodsAgain()
    TestInterfaces()
    TestTypeSwitch()
    TestStringer()
    TestGoroutines()
}

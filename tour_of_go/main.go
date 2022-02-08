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

func main() {
	var i int
	var ii int = 5
	var iii = 5
	iiii := 10
	const str = "Some String"

	fmt.Println("Hello world!")
	fmt.Println("value is ", rand.Intn(10))
	fmt.Println(i, ii, iii, iiii)

    defer fmt.Println("printed last")

	test_loop()
	test_switch()
}

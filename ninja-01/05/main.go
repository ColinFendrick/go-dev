package main

import "fmt"

type foo int

var x foo
var y int

func main() {
	fmt.Printf("x:%v type:%T\n", x, x)
	x = 42
	fmt.Printf("x:%v type:%T\n", x, x)
	y = int(x)
	fmt.Printf("y:%v type:%T", y, y)
}

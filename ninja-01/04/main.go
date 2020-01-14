package main

import "fmt"

type foo int

var x foo

func main() {
	fmt.Printf("x:%v type:%T\n", x, x)
	x = 42
	fmt.Printf("x:%v type:%T", x, x)
}

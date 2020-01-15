package main

import (
	"fmt"
)

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

	{ // Easily scope variables with just some random fucking curly braces in here
		y := 42
		fmt.Println(y)
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

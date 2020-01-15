package main

import "fmt"

func main() {
	a := math(add, 1, 2, 4)
	b := math(mul, 4, 2, 19)
	fmt.Println(a, b)
}

func add(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

func mul(x ...int) int {
	t := 1
	for _, v := range x {
		t *= v
	}
	return t
}

func math(f func(x ...int) int, y ...int) int {
	return f(y...)
}

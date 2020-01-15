package main

import "fmt"

func main() {
	s := foo([]int{1, 2, 3}...)
	m := bar([]int{1, 2, 3, 4})
	fmt.Println(s)
	fmt.Println(m)
}

func foo(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

func bar(x []int) int {
	t := 1
	for _, v := range x {
		t *= v
	}
	return t
}

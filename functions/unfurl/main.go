package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	x := sum(xi...)
	fmt.Println(x)
}

func sum(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

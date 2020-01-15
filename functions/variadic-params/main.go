package main

import "fmt"

func main() {
	str, sum := variadic("Hello", 1, 3, 5, 8, 12308)
	fmt.Println(str, sum)
}

func variadic(a string, x ...int) (string, int) {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return "The starting string is " + a, sum
}

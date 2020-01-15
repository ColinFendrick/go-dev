package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 11}
	res := sum(x...)
	fmt.Println(res)
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

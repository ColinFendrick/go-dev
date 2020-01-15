package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7}
	x = append(x, 8, 9)
	// Spread the x slice into this
	x = append([]int{1, 2, 3}, x...)
	fmt.Println(x)
}

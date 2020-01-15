package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6, 7)

	y := []int{8, 9, 10}
	x = append(x, y...)

	x = append(x[:2], x[3:]...)
	fmt.Println(x)
}

package main

import (
	"fmt"
)

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	_ = append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
}

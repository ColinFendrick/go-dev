package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x[1:])
	fmt.Println(x[:2])

	for i := 0; i <= 3; i++ {
		fmt.Println(x[i])
	}
}

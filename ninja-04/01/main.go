package main

import "fmt"

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	for i, val := range arr {
		fmt.Println(i, val)
	}
	fmt.Printf("%T", arr)
}

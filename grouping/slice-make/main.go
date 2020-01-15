package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println("len:", len(x), "capacity:", cap(x))
	x[9] = 10
	x = append(x, 11, 12)
	fmt.Println(x)
	fmt.Println("len:", len(x), "capacity:", cap(x))
	x = append(x, 13) // Doubles capacity
	fmt.Println(x)
	fmt.Println("len:", len(x), "capacity:", cap(x))
}

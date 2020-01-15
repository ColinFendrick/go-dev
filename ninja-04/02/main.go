package main

import "fmt"

func main() {
	sl := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i <= 9; i++ {
		sl[i] = i
	}
	fmt.Println(sl)
}

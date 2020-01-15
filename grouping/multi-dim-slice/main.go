package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3}
	b := []int{4, 5}
	x := [][][]int{{a, b}, {{6, 7}}}
	fmt.Println(x)
}

package main

import "fmt"

/*
	Pretty much just use slices. Arrays are more for golang internals
*/

func main() {
	arr := [2]int{1, 2}
	multi := [2][2][2]float64{{{3, 4}, {5, 6}}, {{7, 8}, {9, 10}}}
	fmt.Println(arr)
	fmt.Println(multi)
}

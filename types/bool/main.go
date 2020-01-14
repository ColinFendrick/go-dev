package main

import (
	"fmt"
)

var x bool

func main() {
	a := 1
	b := 2
	fmt.Println(x)
	x = true
	fmt.Println(x)
	fmt.Println(b <= a || a == b)
}

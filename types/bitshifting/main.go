package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("kb: %d\t\t\t%b\n", kb, kb)
	fmt.Printf("mb: %d\t\t\t%b\n", mb, mb)
	fmt.Printf("gb: %d\t\t%b\n", gb, gb)
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 3 // 4 times 2, 3 times
	fmt.Printf("%d\t\t%b", y, y)
}

package main

import (
	"fmt"
)

const (
	a int    = 42    // typed
	b        = 42.78 // untyped
	c string = "James Bond"
)

const d = false

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}

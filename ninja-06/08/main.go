package main

import "fmt"

func main() {
	cb := func() func() string {
		return func() string {
			return "Hello world"
		}
	}

	x := cb()
	y := x()
	fmt.Println(y)
	fmt.Println(cb()())
}

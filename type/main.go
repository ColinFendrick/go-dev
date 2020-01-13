package main

import "fmt"

type foo int

func main() {
	var a foo
	a = 42
	b := 400
	c := 100
	// Cannot store b in a because b is type int and a is type foo
	// cannot write a = b or b = a
	// but we can do
	b = int(a)
	// to type-convert a into an int
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("b:%v c:%v\n", b, c)
	c = b
	fmt.Printf("b:%v c:%v", b, c)
}

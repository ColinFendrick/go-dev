/*
A method is nothing more than a FUNC attached to a TYPE.
When you attach a func to a type it is a method of that type.
You attach a func to a type with a RECEIVER.
*/

package main

import "fmt"

type person struct {
	first string
	last  string
}

type number float64

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p1.speak()

	var n number = 12
	fmt.Println("Dividing 12 by two is", n.divByTwo())
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func (n number) divByTwo() number { // Cannot add methods onto non-local types. So do this pattern for primitives
	return n / 2
}

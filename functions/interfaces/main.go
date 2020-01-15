/*
In Go, values can be of more than one type.
An interface allows a value to be of more than one type.
We create an interface using this syntax: “keyword identifier type”
so for an interface it would be: “type human interface”
We then define which method(s) a type must have to implement that interface.
If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}),
then that TYPE implicitly implements the interface and is also of that interface type.
In Go, values can be of more than one type.
*/

package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func main() {
	colin := person{
		first: "Colin",
		last:  "Fendrick",
	}

	bar(colin)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("A person was passed into bar()", h.(person).first)
	default:
		fmt.Println("No corresponding type was found")
	}
}

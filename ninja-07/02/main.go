package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "Colin", last: "Fendrick",
	}

	changeMe(&p)
	fmt.Println(p.first)
}

func changeMe(p *person) {
	p.first = "James"
}

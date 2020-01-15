package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type agent struct {
	person
	live bool
	code string
}

type agentList []agent

func main() {
	james := agent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		live: true,
		code: "007",
	}

	alec := agent{
		person: person{
			first: "Alec",
			last:  "Trevelyan",
			age:   39,
		},
		live: false,
		code: "006",
	}

	colin := person{
		first: "Colin",
		last:  "Fendrick",
		age:   29,
	}

	list := agentList{james, alec}

	fmt.Println(james.person.first)
	fmt.Println(james.last)
	fmt.Println(colin.last)

	for _, a := range list {
		fmt.Println(a.first, a.last, a.code, "alive:", a.live)
	}
}

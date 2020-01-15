package main

import "fmt"

func main() {
	person := struct {
		first string
		last  string
		age   int
	}{
		first: "Colin",
		last:  "Fendrick",
		age:   29,
	}

	table := []struct {
		first string
		last  string
		age   int
	}{
		{
			first: "Colin",
			last:  "Fendrick",
			age:   29,
		},
		{
			first: "James",
			last:  "Bond",
			age:   33,
		},
	}

	fmt.Println(person.first, person.last, person.age)

	for _, p := range table {
		fmt.Println(p.first, p.last, p.age)
	}
}

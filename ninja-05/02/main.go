package main

import "fmt"

type name string
type sList []string

type person struct {
	first              name
	last               name
	favIceCreamFlavors sList
}

func main() {
	colin := person{
		first:              "Colin",
		last:               "Fendrick",
		favIceCreamFlavors: []string{"Peanut Butter", "Salted Carmel", "Coffee"},
	}

	james := person{
		first:              "James",
		last:               "Bond",
		favIceCreamFlavors: []string{"Mint", "Chocolate", "Rum Raisin"},
	}

	people := map[name]person{
		colin.last: colin,
		james.last: james,
	}

	for _, v := range people {
		fmt.Printf("\n%v %v's favorite ice cream flavors, ranked:\n", v.first, v.last)

		for i, v := range v.favIceCreamFlavors {
			fmt.Printf("%d: %v\n", i+1, v)
		}
	}
}

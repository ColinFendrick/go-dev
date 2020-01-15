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

	fmt.Printf("%v %v's favorite ice cream flavors, ranked:\n", colin.first, colin.last)

	for i, v := range colin.favIceCreamFlavors {
		fmt.Printf("%d: %v\n", i+1, v)
	}

	fmt.Printf("%v %v's favorite ice cream flavors, ranked:\n", james.first, james.last)

	for i, v := range james.favIceCreamFlavors {
		fmt.Printf("%d: %v\n", i+1, v)
	}

}

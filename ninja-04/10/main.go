package main

import "fmt"

func main() {
	m := map[string][]string{
		"fendrick_colin": {"DnD", "Yoga"},
		"bond_james":     {"Shaken, not stirred", "Martinis", "Women"},
	}

	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(m, "bond_james")

	for k, v := range m {
		fmt.Printf("%v's favorite things are:\n", k)
		for i, el := range v {
			fmt.Printf("%v\n", el)

			if i == len(v)-1 {
				fmt.Printf("\n")
			}
		}
	}
}

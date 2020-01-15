package main

import "fmt"

func main() {
	m := map[string][]string{
		"fendrick_colin": {"DnD", "Yoga"},
	}

	for k, v := range m {
		fmt.Printf("%v's favorite things are:\n", k)
		for _, el := range v {
			fmt.Printf("%v\n", el)
		}
	}
}

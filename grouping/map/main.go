package main

import "fmt"

// A map is the perfect data structure when you need to look up data fast.

func main() {
	m := make(map[string]string)
	m["hello"] = "world"
	loc := m["hello"]
	fmt.Println(m)
	fmt.Println(loc)
}

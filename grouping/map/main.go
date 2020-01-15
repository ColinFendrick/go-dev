package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["hello"] = "world"
	loc := m["hello"]
	fmt.Println(m)
	fmt.Println(loc)
}

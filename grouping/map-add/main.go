package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	// fmt.Println(m)

	// fmt.Println(m["James"])

	// fmt.Println(m["Barnabas"])

	// v, ok := m["Barnabas"]
	// fmt.Println(v)
	// fmt.Println(ok)

	m["Colin"] = 29

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	for key, val := range m {
		fmt.Println(key, val)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, val := range xi {
		fmt.Println(i, val)
	}
}

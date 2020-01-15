package main

import "fmt"

func main() {
	defer bar("Deferred identifier func call")
	defer func(s string) {
		fmt.Println(s)
	}("Deferred anon func")

	bar("Non-deferred call")
}

func bar(s string) {
	fmt.Println(s)
}

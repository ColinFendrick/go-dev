package main

import "fmt"

func main() {
	s, n := foo(42, "stringy")
	fmt.Println(s)
	fmt.Println(n)
}

func foo(n int, s string) (string, int) {
	return s, n
}

package main

import "fmt"

func main() {
	fmt.Println(func(s string) string {
		return s
	}("Hello, world"))

	func(x int) {
		fmt.Println(x)
	}(42)
}

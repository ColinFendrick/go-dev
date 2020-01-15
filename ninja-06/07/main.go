package main

import "fmt"

func main() {
	x := func(s string) {
		fmt.Println(s)
	}

	x("hello")
}

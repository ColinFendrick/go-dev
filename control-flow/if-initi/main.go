package main

import "fmt"

func main() {
	if x := 42; x == 42 {
		fmt.Printf("x is %v", x)
	}
	// fmt.Printf(x) cannot use x here, it is scoped to the if statement
}

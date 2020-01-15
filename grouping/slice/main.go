package main

import "fmt"

/*
A SLICE holds VALUES of the same TYPE. If I wanted to store all of my favorite numbers, I would use a slice of int.
If I wanted to store all of my favorite foods, I would use a slice of string. We will use a COMPOSITE LITERAL to create a slice.
A composite literal is created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.
*/

func main() {
	sl := []string{"Hello", ""}
	sl[1] = "World"
	sl = append(sl, "!")
	fmt.Println(sl)

	for _, v := range sl {
		fmt.Println(v)
	}
}

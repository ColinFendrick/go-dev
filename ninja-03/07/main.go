package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "false"
	xb, _ := strconv.ParseBool(x)
	if xb {
		fmt.Println("x is true")
	} else if !xb {
		fmt.Println("x is false")
	} else {
		fmt.Println("What?")
		panic("Panicking")
	}
}

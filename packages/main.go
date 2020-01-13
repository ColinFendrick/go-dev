package main

import (
	"fmt"
)

var a string
var b = "string b"

func main() {
	// func Printf(format string, a ...interface{}) (n int, err error)
	n, err := fmt.Println("Hello", 40, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Printed %v btyes \n", n)
	a = "string a"
	fmt.Println(a)
	fmt.Println(b)
}

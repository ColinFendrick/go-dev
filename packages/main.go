package main

import (
	"fmt"
)

func main() {
	// func Printf(format string, a ...interface{}) (n int, err error)
	n, err := fmt.Println("Hello", 40, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Printed %v btyes", n)
}

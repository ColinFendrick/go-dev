package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("true && true: %v\n", true && true)
	fmt.Printf("true && false: %v\n", true && false)
	fmt.Printf("true || true: %v\n", true || true)
	fmt.Printf("true || true: %v\n", true || false)
	fmt.Printf("!true: %v\n", !true)
	a, _ := strconv.ParseBool("true")
	b, _ := strconv.ParseBool("false")
	fmt.Printf("parsebool true || false: %v\n", a && b)
}

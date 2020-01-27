package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Error found",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Printf("foo ran - %v \n", e)
}

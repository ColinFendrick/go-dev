package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			s := fmt.Sprintf("%v is even", i)
			hello(s)
		}
	}
}

func hello(s string) {
	fmt.Println(s)
}

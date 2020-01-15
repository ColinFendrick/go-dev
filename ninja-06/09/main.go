package main

import "fmt"

func main() {
	print := func(s string) {
		fmt.Println(s)
	}
	cb(print, "Hello world")
}

func cb(f func(string), s string) {
	f(s)
}

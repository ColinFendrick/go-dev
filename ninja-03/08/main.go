package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Won't print")
	case true:
		fmt.Println("Will print")
		fallthrough
	case true:
		fmt.Println("Fallthrough print")
	case true:
		fmt.Println("Won't print")
	}
}

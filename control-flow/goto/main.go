package main

import "fmt"

func main() {
	switch {
	case true:
		if true {
			goto nextCase
		}
		break
	nextCase:
		fallthrough
	default:
		fmt.Println("Finally")
	}
}

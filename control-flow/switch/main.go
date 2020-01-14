package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough // Goes to next case
	case (4 == 4):
		fmt.Println("also true, does it print?") // This will print if fallthrough is used above
	case true:
		fmt.Println("Won't print, since there's no fallthrough above")
	}
}

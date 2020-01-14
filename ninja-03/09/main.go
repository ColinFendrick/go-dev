package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "Basketball":
		fmt.Println("It's case sensitive!")
	case "basketball":
		fmt.Println("This will print")
		fallthrough
	case "baseball":
		fmt.Println("This will print even though the condition is false cause of the fallthrough")
		fallthrough
	default:
		fmt.Println("Got to default due to fallthrough")
	}
}

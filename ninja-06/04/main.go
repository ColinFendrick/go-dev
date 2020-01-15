package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	colin := person{
		first: "Colin",
		last:  "F",
		age:   29,
	}

	fmt.Println(colin.speak("Hello I'm "))
}

func (p person) speak(s string) string {
	return s + p.first + " " + p.last
}

package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak(string)
}

func main() {
	p := person{
		name: "Colin Fendrick",
	}

	saySomething(&p)
	// saySomething(p) <-- This fails
}

func (p *person) speak(s string) {
	fmt.Println(s)
}

func saySomething(h human) {
	h.speak("Hello")
}

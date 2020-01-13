package main

import (
	"fmt"
)

var a int
var b string = "James Bond"
var c bool
var d bool = true

func main() {
	e := 42
	f := `Shaken not stirred`

	fmt.Println(a)
	fmt.Printf("%v\n", b)
	fmt.Print(c, "\n")

	fmt.Sprintln(d)
	i := fmt.Sprintf("%v\n", e)
	fmt.Print(i)
	j := fmt.Sprint(b, " says, ", f)
	fmt.Print(j)
}

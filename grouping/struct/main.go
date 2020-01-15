package main

import "fmt"

type vals struct {
	x, y int
	u    float32
}

func main() {
	s := vals{x: 1, y: 4, u: 3.14}
	fmt.Println(s.u)
}

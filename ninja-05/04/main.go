package main

import "fmt"

func main() {
	car := struct {
		doors int
		color string
	}{
		doors: 4,
		color: "Grey",
	}

	fmt.Println(car.doors, car.color)
}

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Printf("Truck is %v with %d doors and four wheel drive is %v", t.vehicle.color, t.doors, t.fourWheel)
	fmt.Printf("Sedan is %v with %d doors and luxury = %v", s.color, s.doors, s.luxury)
}

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
	ram := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	q50 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		luxury: true,
	}

	fmt.Println(ram)
	fmt.Println(ram.color)
	fmt.Println(q50)
	fmt.Println(q50.luxury)

}

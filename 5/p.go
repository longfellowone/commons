package main

import "fmt"

type cars struct {
	name     string
	topSpeed int
	features []string
}

func main() {
	lot := []*cars{}

	car := new(cars)
	car.name = "Lancer"
	car.topSpeed = 250
	car.features = []string{"AC", "Power Windows"}

	lot = append(lot, car)

	car = new(cars)
	car.name = "Ram"
	car.topSpeed = 160
	car.features = []string{"Sun roof", "Audio"}

	lot = append(lot, car)

	fmt.Println("--------")

	for i := range lot {
		fmt.Println(lot[i])
	}

	fmt.Println("--------")

	for _, v := range lot {
		fmt.Println(v.name)
		for _, v := range v.features {
			fmt.Println(v)
		}
	}

}

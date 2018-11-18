package main

import "fmt"

type persons struct {
	id        int
	fname     string
	age       int
	ownsHouse bool
	favCar    []string
}

func main() {

	p2 := persons{
		id:        2,
		fname:     "Vivian",
		age:       29,
		ownsHouse: false,
		favCar:    []string{"Lancer", "Lambo", "Porsche"},
	}

	p3 := persons{
		id:        3,
		fname:     "Ken",
		age:       60,
		ownsHouse: true,
		favCar:    []string{"Camery", "Corolla", "Ram"},
	}

	p := []persons{p2, p3}

	for _, v := range p {
		fmt.Println(v.id, v.fname, v.age, v.ownsHouse)
		for _, v := range v.favCar {
			fmt.Println(v)
		}
	}
	//p := struct {
	//	fname string
	//	lname string
	//	age   int
	//}{
	//	fname: "Matt",
	//	lname: "Wright",
	//	age:   28,
	//}
	//
	//fmt.Println(p.fname, p.lname, p.age)

}

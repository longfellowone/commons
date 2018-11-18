package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIcecream []string
}

func main() {
	p1 := person{
		firstName: "Matt",
		lastName:  "Wright",
		favIcecream: []string{
			"Vanilla",
			"Chocolate",
		},
	}

	p2 := person{
		firstName: "Vivian",
		lastName:  "Moser",
		favIcecream: []string{
			"Haznelnut",
			"Bummble gum",
		},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for i := range p1.favIcecream {
		fmt.Println(p1.favIcecream[i])
	}

	fmt.Println("\n")

	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.favIcecream {
		fmt.Println(v)
	}
}

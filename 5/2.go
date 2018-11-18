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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for _, v := range v.favIcecream {
			fmt.Println(v)
		}
	}
}

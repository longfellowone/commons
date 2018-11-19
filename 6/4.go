package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return fmt.Sprintf("My name is %v %v and I am %d years old", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Matt",
		last:  "Wright",
		age:   28,
	}

	fmt.Println(p.speak())
}

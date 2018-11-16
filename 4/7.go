package main

import "fmt"

func main() {
	b := []string{"James", "Bond", "Shaken, not stirred"}
	m := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xxx := [][]string{b, m}

	for i, v := range xxx {
		fmt.Println("------")
		fmt.Println("Index: ", i)
		for _, n := range v {
			fmt.Println(" ", n)
		}
	}

	fmt.Println("------")
}

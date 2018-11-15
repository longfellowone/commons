package main

import "fmt"

func main() {
	favSport := "baseball"
	switch favSport {
	case "hockey":
		fmt.Println("Sport is hockey")
	case "baseball":
		fmt.Println("Sport is baseball")
	default:
		fmt.Println("You did not choose a valid sport!")
	}
}

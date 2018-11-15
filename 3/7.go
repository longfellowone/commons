package main

import "fmt"

func main() {
	x := "car"
	if x == "train" {
		fmt.Println("X is train")
	} else if x == "bike" {
		fmt.Println("X is bike")
	} else {
		fmt.Println("X is neither train nor bike")
	}
}

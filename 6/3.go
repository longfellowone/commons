package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("First")
	}()

	fmt.Println("Second")
}

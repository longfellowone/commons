package main

import "fmt"

func main() {
	xx := []int{1, 2, 3, 4, 5}

	for _, v := range xx {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", xx)
}

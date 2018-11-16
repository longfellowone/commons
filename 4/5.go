package main

import "fmt"

func main() {
	xx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xx = append(xx[:3], xx[6:]...)
	fmt.Println(xx)
}

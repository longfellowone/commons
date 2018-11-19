package main

import (
	"fmt"
)

func main() {

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	foo := foo(n...)

	fmt.Println("Sum of foo", foo)

	fmt.Println("Sum of bar", bar(n))

}

func foo(nn ...int) int {
	t := 0
	for _, n := range nn {
		t += n
	}
	return t
}

func bar(nn []int) int {
	t := 0
	for _, n := range nn {
		t += n
	}
	return t
}

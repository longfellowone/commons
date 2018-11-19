package main

import (
	"fmt"
	"os"
)

func sum(x ...int) (t int, err error) {
	if x == nil {
		err = fmt.Errorf("no numbers to sum")
		return
	}

	if len(x) == 1 {
		err = fmt.Errorf("provide more than one number")
		return
	}

	for _, v := range x {
		fmt.Printf("Adding %d to %d for a total of %d\n", v, t, v+t)
		t += v
	}
	return
}

func even(sum func(x ...int) (int, error), xx ...int) (int, error) {
	var t []int
	for _, v := range xx {
		if v%2 == 0 {
			t = append(t, v)
		}
	}
	if t == nil {
		return 0, fmt.Errorf("no even numbers to sum")
	}
	if len(t) == 1 {
		return 0, fmt.Errorf("provide more than one even number")
	}
	return sum(t...)
}

func main() {
	nums := []int{3, 3, 4, 8}

	total, err := sum(nums...)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}

	fmt.Println("Total is:", total, "\n")

	total, err = even(sum, nums...)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}

	fmt.Println("Total of even numbers is:", total)
}

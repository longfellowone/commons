package main

import "fmt"

func main() {

	m := map[string]int{
		"vivian": 28,
		"matt":   29,
		"diana":  32,
	}

	fmt.Println(m, "\n")

	for k, v := range m {
		fmt.Printf("%v's age is %d\n", k, v)
	}

	//fmt.Println("\n")
	//delete(m, "matt")
	//
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}

}

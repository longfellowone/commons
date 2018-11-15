package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for n := 1; n <= 3; n++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

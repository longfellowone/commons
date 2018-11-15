package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "\t", i%4)
	}
}

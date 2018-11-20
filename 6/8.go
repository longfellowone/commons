package main

import "fmt"

func printer(msg string) func() string {
	return func() string {
		return fmt.Sprint("Printing... ", msg)
	}
}

func main() {
	msg := printer("Hello World")

	fmt.Println(msg())
}

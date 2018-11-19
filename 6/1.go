package main

import "fmt"

func main() {

	f := foo()
	b, s := bar()

	fmt.Println(f, b, s)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "text"
}

package main

import "fmt"

func main() {
	msg := func() string { return fmt.Sprint("My message") }
	fmt.Println(msg())
}

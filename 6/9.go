package main

import "fmt"

func printer(f func(msg string) string, w string) string {
	return fmt.Sprint(f("Printing..."), "...", w)
}

func main() {

	msg := func(msg string) string {
		return fmt.Sprint(msg, "My message")
	}
	when := "now!"

	fmt.Println(printer(msg, when))
}

package main

const current = 2018

const (
	a = current - iota // Count down backwards
	b
	c
	d
)

func main() {
	println(a, b, c, d)
}

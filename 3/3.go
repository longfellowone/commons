package main

import "fmt"

//const (
//	birthday    = 1990
//	currentYear = 2018
//)

func main() {
	bd := 1990

	for i := 0; bd <= 2018; i++ {
		fmt.Println(i)
		bd++
	}

	//between := currentYear - birthday
	//
	//for i := 1; i <= between; i++ {
	//	fmt.Println(i)
	//}
}

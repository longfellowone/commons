package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice creame", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println("-----")
		fmt.Println("Name:", k)
		for _, v := range v {
			fmt.Println(v)
		}
	}
}

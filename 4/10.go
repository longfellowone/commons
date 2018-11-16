package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice creame", "Sunsets"},
	}
	m["matt"] = []string{"cars", "pizza", "computers"}

	for k, v := range m {
		fmt.Println(k, "likes...")
		for _, v := range v {
			fmt.Println(v)
		}
	}

	delete(m, "matt")
	delete(m, "bond_james")
	delete(m, "no_dr")

	fmt.Printf("\nDELETED 3 RECORDS\n\n")
	for k, v := range m {
		fmt.Println(k, "likes...")
		for _, v := range v {
			fmt.Println(v)
		}
	}
}

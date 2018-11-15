package main

import (
	"fmt"
	"os"
)

// removeInt returns new slice with value removed
func removeInt(f int, s []int) ([]int, error) {
	for i, v := range s {
		if f == v {
			r := append(s[:i], s[i+1:]...)
			t := make([]int, len(s)-1)
			// Another way to delete from slice
			//x[2] = x[len(x)-1] // Sent x to end of index
			//x = x[:len(x)-1]   // Chop off x from end of index
			// And another way with sort.IntSlice.Search
			copy(t, r)
			return t, nil
		}
	}
	return s, fmt.Errorf("did not find %d, unable to remove", f)
}

func main() {
	x := []int{44, 55, 66, 34, 21, 50}
	f := 65

	d, err := removeInt(f, x)

	if err != nil {
		fmt.Printf("Error: %v\n%v\n", err, d)
		os.Exit(0)
	}
	fmt.Println(d)

}

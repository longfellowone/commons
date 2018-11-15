package _

import "fmt"

func main() {
	i := 0
	for {
		if i > 28 {
			break
		}
		fmt.Println(i)
		i++
	}
}

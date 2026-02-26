package main

import (
	"fmt"
)

func main() {

	var fruits map[string]int
	fruits = map[string]int{"Apple": 5, "Banana": 10}
	// add orange
	fruits["orange"] = 15
	// show all fruits
	for fruit, quantatiy := range fruits {
		fmt.Printf("%s,%d\n", fruit, quantatiy)
	}
}

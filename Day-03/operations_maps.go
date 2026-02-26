package main

import (
	"fmt"
)

func main() {
	fmt.Println("Main fun started")
	m := map[string]int{"Apple": 2, "banana": 4}
	fmt.Println(len(m))
	m["Mango"] = 6
	fmt.Println("Mango added", m)
	// delete(m, "banana")
	fmt.Println("banana deleted", m)
	banana_value := m["banana"]
	println(banana_value)
}

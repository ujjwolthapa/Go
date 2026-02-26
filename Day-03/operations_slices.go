package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Main func started")
	numbers := make([]int, 10, 20)
	numbers = append(numbers[:0], 1)
	numbers = append(numbers, 3, 150, 2, 3, 4, 100, 3)
	fmt.Println(numbers)
	fmt.Printf("len %d and cap %d", len(numbers), cap(numbers))
	// num := numbers[:2]
	// fmt.Println(num)
	numbers = append(numbers[:1], numbers[2:]...)
	fmt.Println(numbers)

	if len(numbers) == 0 {
		fmt.Print("Empty\n")
	} else {
		fmt.Print("Not empty\n")
	}
	for i, v := range numbers {
		if v == 100 {
			fmt.Printf("100 found\n")
			fmt.Printf("Found on %d\n", i)
		}
	}
	// sorted numbers
	sort.Ints(numbers)
	fmt.Println("sorted", numbers)
}

package main

import "fmt"

func main() {
	r := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println("r", r)
	fmt.Println("Alice=", r["Alice"])
	fmt.Println("Bob=", r["Bob"])
}

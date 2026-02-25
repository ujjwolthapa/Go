package main

import (
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	var q Person = Person{"Arzoo", 26}
	fmt.Println("this is struct", q)
	fmt.Println("name=", q.Name)
	fmt.Println("age=", q.Age)
}

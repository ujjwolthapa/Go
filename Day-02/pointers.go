package main

import (
	"fmt"
)

func main() {
	var s int = 10
	var t *int = &s
	fmt.Println("this is value", s)
	fmt.Println("this is pointer value", t)
	fmt.Println("print value at memory address stored in t", *t)
}

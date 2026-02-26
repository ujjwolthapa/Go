package main

import (
	"fmt"
)

func main() {
	var ptr1 *int
	fmt.Println("ptr1", ptr1)
	x := 1
	ptr2 := &x
	fmt.Println("ptr2", *ptr2)
	ptr3 := new(int)
	*ptr3 = 300
	fmt.Println("ptr3", ptr3)
	ptr4 := &ptr3
	fmt.Println("ptr4", **ptr4)

}

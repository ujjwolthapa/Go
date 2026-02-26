package main

import (
	"fmt"
)

func increment(x *int) {
	*x++
	fmt.Println("Increment from function", *x)
}
func main() {
	var x int = 10
	// var p *int = &x
	p := &x
	fmt.Println("value of mem add", p)
	fmt.Println("value of x", *p)
	*p = 20
	// after value change
	fmt.Println("value is changed", p)
	var y int = 22
	// p1 := &y
	fmt.Println(&x)
	increment(&x)
	increment(&y)

}

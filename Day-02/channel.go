package main

import "fmt"

func main() {
	fmt.Println("channel")
	u := make(chan int)
	fmt.Println("u", u)
}

package main

import (
	"fmt"
)

func main() {
	var arr [3][3]int
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = count
			count++
		}

	}
	// print
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d", arr[i][j])
		}
		println()
	}

}

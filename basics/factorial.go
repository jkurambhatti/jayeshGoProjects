package main

import (
	"fmt"
)

func main() {

	fmt.Println(factorial(4))
}

func factorial(x int) int {

	if x == 1 {
		return 1
	} else {
		return (x * factorial(x-1))
	}
}

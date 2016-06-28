package main

import (
	"fmt"
)

func rotate(a []int, n int) {
	temp := make([]int, 4)
	copy(temp[:n], a[len(a)-n:])
	copy(temp[n:], a[:n])
	copy(a, temp)
	// fmt.Println(a)
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	rotate(arr, 3)
	fmt.Println(arr)
}

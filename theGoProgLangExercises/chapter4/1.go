package main

import (
	"crypto/sha256"
	"fmt"
)

var tb [8]byte

func init() {
	for i := uint(0); i < 8; i++ {
		tb[i] = byte(1 << i)
		fmt.Println("%b", tb)
	}
}

func countDifferent(n1, n2 [32]byte) int {
	var count int

	for i := 0; i < 32; i++ {
		h1 := n1[i]
		h2 := n2[i]

		for j := 0; j < 8; j++ {
			if h1&tb[j] != h2&tb[j] {
				count++
			}
		}
	}
	return count
}

func main() {

	has1 := sha256.Sum256([]byte("a"))
	has2 := sha256.Sum256([]byte("b"))

	fmt.Printf("%x\n%x\n", has1, has2)

	n := countDifferent(has1, has2)
	fmt.Println(n)
}

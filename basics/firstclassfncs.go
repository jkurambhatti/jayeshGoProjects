package main

import "fmt"

type interff interface{}

func call() interff {
	return func(args ...int) int {
		fmt.Println("hello fif")
		sum := 0
		for _, val := range args {
			sum += val
		}
		return sum
	}
}

func main() {
	fun := call()
	switch x := fun.(type) {
	case func(...int) int:
		fmt.Println(x(10, 20, 30, 40))
	}

}

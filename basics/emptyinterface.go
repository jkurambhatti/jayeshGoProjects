// we will see what an empty interface is where is it used

// this example shows how we can pass anything to a function that implements an empty interface

// empty interface is empty, so it can receive anything

package main

import (
	"fmt"
)

type printer interface{}

func do_print(p printer) {
	fmt.Println(p)
}

func main() {
	do_print(2)
	do_print(1.15)
	do_print("hello world")

}

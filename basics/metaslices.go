// we will see here how to create an array inside an array or slice inside another slice

// op : [1,3,5,[7,9,11]]

package main

import (
	"fmt"
)

type mslicer interface{}

var (
	myslice = make([]mslicer, 0) // array/slice of an empty interface, which means it can hold anything
)

func appendiser(m mslicer) {
	fmt.Printf("type of m %T \n", m)
	//fmt.Println(m.(string))
	myslice = append(myslice, m)

}

func main() {
	appendiser(10)
	appendiser(15.2568)
	appendiser("hello")
	//a := []int{1, 2, 3, 4, 5}
	//b := []string{"ab", "bc", "cd", "ef", "gh"}
	//appendiser(a)
	//appendiser(b)
	fmt.Println(myslice)
	fmt.Println(myslice[0])

}

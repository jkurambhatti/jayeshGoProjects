// how to create functions in golang
// closures are functions defined inside functions that have access to local variables

package main 

import "fmt"

func main() {
//	result := add(10,20)
a,b := nextTwo(10,20)    // returns 2 values

subtract := func() {
	a--
	b--
}

	subtract()
	subtract()

 fmt.Println(sumup(10,20,30,40))


	fmt.Println(a,b)
}

func nextTwo(x,y int) (int, int) {
	return x+1, y+1
	
}

func add(a,b int) int{
	return a+b
}


// variable arguments functions
func sumup(args ... int)  int {
	sum := 0

	for _, value := range args {
		sum = sum + value
	}

	return sum

}


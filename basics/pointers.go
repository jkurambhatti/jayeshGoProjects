package main 

import "fmt"

func main() {
	 y := 0
	 changeMyVal(y)
	 fmt.Println(y)
	 changeMyValPoint(&y)
	 fmt.Println(y)

}

func changeMyVal(x int) {
	x = 100	
}

func changeMyValPoint(x *int) {
	*x = 100
}
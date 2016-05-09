// defer keyword makes the function call after it to execute after finishing the complete function

package main 
import "fmt"

func main() {
	
//	simpleDiv(5,0)
 	  printOne()
 	  panic("PANIC !!!!")
	  printTwo()


}

func printOne() {
	fmt.Println(1)
}


func printTwo() {
	fmt.Println(2)
}

func simpleDiv(a,b int) int {

	defer recover() 

	return a/b
}






package main

import "fmt"

func main() {
	/*
	   var myname string = "jayesh"
	   var lastname = "kurambhatti"
	   var age = 23
	   var is100 bool = true


	   fmt.Println(myname,lastname,age)
	   fmt.Println(len(myname))
	   fmt.Println(myname + "is a genius")
	   fmt.Printf("%T\t%T\n",myname,age )
	   fmt.Printf("%T",is100)
	   fmt.Printf("%t",is100)

	   // logical operators &&, ||, !

	   fmt.Println("true && false\n",true && false)

	   fmt.Println("true || false\n",true || false)

	   fmt.Println("!true \n",!true)
	*/

	// for loops
	/*
	   i := 1

	   for i<=5 {
	    fmt.Println(i)
	    i++
	   }

	   for j:=10; j>0; j-- {
	     fmt.Println(j)
	   }



	*/

	// conditions
	//  x ,y := 10, 20
	x := 10
	/*
	    if x>y {
	   		fmt.Println("x true")
	   	}  else if x<y {
	   		fmt.Println("x false")
	   	} else {
	   		fmt.Println("equal")
	   	}
	*/
	switch x {
	case 10:
		fmt.Println("x true")
	case 20:
		fmt.Println("x false")
	default:
		fmt.Println("equal")

	}

}

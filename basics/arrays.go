package main

import "fmt"

func main() {

	// decalaration
	/*
	   var marks[5] int

	   fmt.Println(marks[3])   // all elements are initialized to zero

	   //definition
	*/
	arr := []string{"ab", "cd", "ef", "gh", "ij", "kl"}

	// print all values using range
	/*
	   for _ , x := range arr {   // first par : index, 2nd : avctual value at index
	   fmt.Println(x)
	   }
	*/

	// slicing [x:y]

	fmt.Println(arr[:])  // prints everything
	fmt.Println(arr[:3]) // prints beg to 2, and 3 not included
	fmt.Println(arr[2:]) // prints from 2 till end

	ar1 := arr[:] // shallow copy copies reference also
	// anychange in either affects the other one

	fmt.Println("ar1 ", ar1[:]) // prints everything

	ar1[2] = "xz"

	//fmt.Println("ar1[2] ",ar1[2])  // prints everything
	//fmt.Println("arr[2] ",arr[2])  // prints everything

//	fmt.Println("ar1 ", ar1[:]) // prints everything
//	fmt.Println("arr ", arr[:]) // prints everything

	//deep copy using make() and copy()

	ar2 := make([]string,5)
	copy(ar2,arr)

	fmt.Println("ar2",ar2) // prints everything
//	ar2[2] = "12"		    // did not affect the arr[]	
	ar2 = append(ar2,"45")
	ar2 = append(ar2,"78","99")
	fmt.Println("ar2",ar2) // prints everything
//	fmt.Println("ar2", ar2[:]) // prints everything
	for i, val := range ar2{
	fmt.Println(i,val);

	fmt.Println(ar2 + 2)
	}


}

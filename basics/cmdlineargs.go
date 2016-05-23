package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("cmdline args are :")
	for i, _ := range os.Args {
		fmt.Printf("%d : %s \n", i, os.Args[i])
	}
}

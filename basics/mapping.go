package main 

import (
	"fmt"
)

type variable map[string]string

type env struct{
	v variable	
}

var op = variable{
	"add" : "do addition",
}

func solveex(e *env, y string) {

	fmt.Println(e.v[y])
}

func main()	
	solveex(&op, "add")
}

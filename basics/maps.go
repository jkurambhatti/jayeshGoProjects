package  main
	
import "fmt"

func main() {
	// how to create maps 
	// maps is a collection of key:value pairs

	nameage := make(map[string] int)

	nameage["jayesh"] = 23
	nameage["akash"] = 24

	fmt.Println(nameage["jayesh"])
	fmt.Println(len(nameage))
}

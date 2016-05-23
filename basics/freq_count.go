package main

import (
	"fmt"
	//	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println("error")
	}
}

var (
	str        string
	stop_words = []string{"is", "an", "the", "a", "on", "if", "this", "that", "it", "for", "be", "to", "all", "and", "in", "of", "as", "or"}
	table      = make(map[string]int)
)

func main() {
	fp, err := os.Open("file.txt")
	check(err)

	buf := make([]byte, 1024)

	for {
		n, err := fp.Read(buf)
		if n == 0 {
			fmt.Println("file reading finished")
			check(err)
			break
		} else {

			count := 0
			str := string(buf)

			//split based on " "
			list := strings.Fields(str)

			// remove non alphabet characters
			for k := 0; k < len(list); k++ {

				// convert string to lower case
				list[k] = strings.ToLower(list[k])

				if strings.Contains(list[k], ".") {
					list[k] = strings.Replace(list[k], ".", "", -1)
				} else if strings.Contains(list[k], ";") {
					list[k] = strings.Replace(list[k], ";", "", -1)
				} else if strings.Contains(list[k], ",") {
					list[k] = strings.Replace(list[k], ",", "", -1)
				}
			}

			// remove stop words
			for i := 0; i < len(list); i++ {
				for j := 0; j < len(stop_words); j++ {
					if list[i] == stop_words[j] {
						list = append(list[:i], list[i+1:]...)
					}
				}
			}

			// delete duplicate elements
			// count occurences of elements

			for i := 0; i < len(list); i++ {
				for j := i; j < len(list); j++ {
					if list[i] == list[j] {
						count++
						if count > 1 {
							list = append(list[:j], list[j+1:]...)
						}
					}
				}
				//fmt.Println(list[i], count)
				table[list[i]] = count // map of string and count
				count = 0
			}

			fmt.Println("table :  \n", table)
			// sort the elements

		}

	}

}

package main

import (
    "io"
    "os"
)

func main() {
    // open input file
    fi, err := os.Open("date.go")
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    // open output file
    fo, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := fi.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        // write a chunk
        if _, err := fo.Write(buf[:n]); err != nil {
            panic(err)
        }
    }
}



/*

==============================================================================================================
package main

import (
	//"fmt"
	//"io"
	//"io/ioutil"
	//"log"
	"os"
)

func check_err(err error) {
	panic(err)
}

func main() {

	_, err := os.Open("/home/jayesh/workspace/go/src/jayeshGoProjects/basics/date.go") // open a file
	check_err(err)

	b1 := make([]byte, 500)
	n1, err := fp.Read(b1)
        check_err(err)
        fmt.Printf("%d bytes: %s\n", n1, string(b1))


	file_info, err := os.Stat("date.go")
	fmt.Println(file_info.Size())
	b1 := make([]byte, 50)
	n, err := fp.Read(b1)
	fmt.Printf(" %d bytes read ", n, string(b1))


*/

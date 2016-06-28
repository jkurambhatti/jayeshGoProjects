// Exercis e 4.2: Wr ite a program that prints the SHA256 hash of its stand ard inp ut by defau lt but
// supp orts a command-line flag to print the SHA384 or SHA512 hash ins tead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s256, s384, s512 *bool

func init() {
	s256 = flag.Bool("sha256", false, "do sha256 hashing")
	s384 = flag.Bool("sha384", false, "do sha384 hashing")
	s512 = flag.Bool("sha512", false, "do sha512 hashing")
}

func main() {
	// args := os.Args
	flag.Parse()
	fmt.Println(os.Args[2])
	if *s384 {
		fmt.Println("s384", *s384)
		h1 := sha512.Sum384([]byte(os.Args[2]))
		fmt.Printf("%x\n", h1)
	} else if *s512 {
		fmt.Println("s512", *s512)
		h1 := sha512.Sum512([]byte(os.Args[2]))
		fmt.Printf("%x\n", h1)
	} else {
		fmt.Println("s256", *s256)
		h1 := sha256.Sum256([]byte(os.Args[2]))
		fmt.Printf("%x\n", h1)
	}

}

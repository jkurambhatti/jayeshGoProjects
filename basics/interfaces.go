// demonstration of interfaces in golang and their use cases

package main

import (
	"fmt"
)

type inter interface {
	area() int
	perim() int
}

type rect struct {
	width, height int
}

type circle struct {
	radius int
}

type triangle struct {
	base, height int
}

// define methods associated with struct

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return (r.width + r.height) * 2
}

func (t triangle) area() int {
	return int(t.base * t.height)
}

func (t triangle) perim() int {
	return (t.base + t.height) * 2
}

func (c circle) perim() int {
	return int(3*c.radius) * 2
}

func (c circle) area() int {
	return int(3 * c.radius * c.radius)
}

func calculate(shape inter) int {
	//fmt.Printf(" type of shape : %T , area = %d\n", shape, shape.area())
	//fmt.Printf(" type of shape : %T , perim = %d\n", shape, shape.perim())
	a := shape.area()
	p := shape.perim()

	fmt.Println(a, p)
	return 0
}

func main() {
	var r1 = rect{10, 20}
	var c1 = circle{100}
	var t1 = triangle{10, 20}

	fmt.Println(calculate(r1))
	fmt.Println(calculate(c1))
	fmt.Println(calculate(t1))
}

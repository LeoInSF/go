package main

import "fmt"

type rect struct {
	width, height int
}

//methods defined on struct types.

// pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// value receiver
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	// Go automatically handles conversion between values and pointers for method calls.
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}

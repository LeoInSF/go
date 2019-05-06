package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 5
	for {
		fmt.Println("Do Stuff,", x)
		x += 3
		if x > 25 {
			break
		}
	}
}

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// fmt.Println("Welcome to go world!")
	// foo()
	// randNum()
	// var num1 = 5.6
	// var num2 = 9.5
	// fmt.Println(add(num1, num2))
	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
}

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func randNum() {
	fmt.Println("A number from 1-100", rand.Intn(100))
}

func add(x float64, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

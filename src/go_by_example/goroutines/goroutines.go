package main

// A goroutine is a lightweight thread of execution.
import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("new")
	fmt.Scanln()
	fmt.Println("done")
}

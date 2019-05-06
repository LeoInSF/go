package main

import "fmt"

func main() {
	// var grades map[string]float32
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67
	fmt.Println(grades)

	TimesGrade := grades["Timmy"]
	fmt.Println(TimesGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)
	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
	// _, score := range grades
}

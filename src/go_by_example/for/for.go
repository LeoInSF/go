package main

import  "fmt"

func main()  {
  // Baic for loop with single condition
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }
  // a classic initial/condition/after for loop
  for j := 7; j <= 9; j ++ {
    fmt.Println(j)
  }

  // for without a condition will loop repeatedly until you break out of the llop or return from the enclosing function
  for {
    fmt.Println("loop")
    break
  }

  // you can also continue to the next iteration of the loop
  for n := 0; n <= 5; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}

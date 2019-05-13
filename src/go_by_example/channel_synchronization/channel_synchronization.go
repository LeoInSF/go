package main

/*
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
*/

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working ...")
	time.Sleep(5 * time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
	// Block until we receive a notification from the worker on the channel.
	go worker(done)
	<-done
}

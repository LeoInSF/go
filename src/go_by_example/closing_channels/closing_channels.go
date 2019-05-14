package main

/*
Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel's receivers.
*/
import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	/*
	  It repeatedly receives from jobs with j, more := <-jobs.
	  In this special 2-value form of receive,
	  the more value will be false if jobs has been closed and all values in the channel have already been received.
	  We use this to notify on done when we’ve worked all our jobs.
	*/
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// We await the worker using the synchronization approach we saw earlier.
	<-done
}

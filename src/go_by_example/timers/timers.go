package main

/*
We often want to execute Go code at some point in the future, or repeatedly at some interval.
Go’s built-in timer and ticker features make both of these tasks easy.
*/
import "time"
import "fmt"

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	// One reason a timer may be useful is that you can cancel the timer before it expires.
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// The first timer will expire ~2s after we start the program,
	// but the second should be stopped before it has a chance to expire.
}

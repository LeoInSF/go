package main

/*
Another option is to use the built-in synchronization features of goroutines and channels to achieve the same
result. This channel-based approach aligns with Go's ideas of sharing memory by communicating and having
each piece of data by exactly 1 goroutine
*/
import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/* state will be owned by a single goroutine. This will guarantee that the data is never corrupted
with concurrent access. In order to read or write that state, other goroutines will send messages
to the owning goroutine and receive corresponding replies.
*/

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64
	// The reads and writes channels will be used by other goroutines to issue read and write requests, respectively
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// NOTE: This goroutine repeatedly selects on the reads and writes channels,
	//responding to requests as they arrive. A response is executed by first performing
	//the requested operation and then sending a value on the response channel resp to indicate
	//success (and the desired value in the case of reads).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				write.resp <- true
			}
		}
	}()
	// This  starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel.
	// Each read requires constructing a readOp, sending it over the reads channel,
	// and the receiving the result over the provided resp channel.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	/*
	   For this particular case the goroutine-based approach was a bit more involved than the mutex-based one.
	   It might be useful in certain cases though, for example where you have other channels involved or
	   when managing multiple such mutexes would be error-prone.
	   You should use whichever approach feels most natural,
	   especially with respect to understanding the correctness of your program.
	*/

}

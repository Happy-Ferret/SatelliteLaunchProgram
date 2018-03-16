package main

import (
	"fmt"
	"sync"
	"time"
)

type satsHolder struct {
	wg           sync.WaitGroup
	satsLaunched int
	totalSats    int
}

func main() {

	LaunchStation1 := make(chan bool)
	LaunchStation2 := make(chan bool)

	h := satsHolder{
		satsLaunched: 0,
		totalSats:    500,
	}

	h.wg.Add(2) // Wait for the goroutines to finish before

	// Read from LaunchStation2, Write to LaunchStation1
	go LS1(&h, LaunchStation2, LaunchStation1)
	// Read from LaunchStation1, Write to LaunchStation2
	go LS2(&h, LaunchStation1, LaunchStation2)

	LaunchStation2 <- true // sends a message to LS1 to start

	h.wg.Wait()         // Wait until goroutines have finished
	fmt.Println("Done") // Done launching
}

// LS1 takes a satsHolder, a channel to readFrom (LaunchStation2), a channel to writeTo (LaunchStation1)
// LS1 waits until it gets a message on readFrom then launches the next 4 satellites
// Once the satellites are launched it sends a message on writeTo
// The cycle continues until all satellites are launched
func LS1(sat *satsHolder, readFrom <-chan bool, writeTo chan<- bool) {
	defer close(writeTo) // Close channel LaunchStation1 when goroutine is completed
	defer sat.wg.Done()  // Tells wg (WaitGroup) that it has finished

	for {
		select {
		case <-time.After(5 * time.Second): // 5 second timeout
			return
		case <-readFrom:
			if sat.satsLaunched < sat.totalSats {
				fmt.Println("Launching from LS1")
				for i := 0; i < 4; i++ {
					sat.satsLaunched++
					fmt.Printf("3, 2, 1... Launched SAT%d\n", sat.satsLaunched)

				}
				writeTo <- true // send message that launches were successful
			} else {
				return // no more sats to launch
			}
		}
	}
}

// LS2 takes a satsHolder, a channel to readFrom (LaunchStation1), a channel to writeTo (LaunchStation2)
// LS2 waits until it gets a message on readFrom then Launches the next 4 satellites
// Once the satellites are launched it sends a message on writeTo
// The cycle continues until all satellites are launched
func LS2(sat *satsHolder, readFrom <-chan bool, writeTo chan<- bool) {
	defer close(writeTo) // Close channel LaunchStation2 when goroutine is completed
	defer sat.wg.Done()  // Tells wg (WaitGroup) that it has finished

	for {
		select {
		case <-time.After(5 * time.Second): // 5 second timeout
			return
		case <-readFrom: // Waits until it gets a message from LS1
			if sat.satsLaunched < sat.totalSats {
				fmt.Println("Launching from LS2")
				for i := 0; i < 4; i++ {
					sat.satsLaunched++
					fmt.Printf("3, 2, 1... Launched SAT%d\n", sat.satsLaunched)
				}
				writeTo <- true // sends message that launches were successful (ok for LS1 to start)
			} else {
				return // no more sats to launch
			}
		}
	}
}

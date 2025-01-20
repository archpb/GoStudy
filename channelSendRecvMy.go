package main

import (
	"fmt"
	"sync"
)

func main() {
	// test server routine send data to one channel, main receive use slice
	mainChan := make(chan int, 4)
	var wg sync.WaitGroup
	var wgRecv sync.WaitGroup // add sync for receive route

	// Start receiving data concurrently
	recv := make([]int, 0, 10) // Preallocate space for up to 30 elements
	wgRecv.Add(1)
	go func() {
		defer wgRecv.Done()
		for v := range mainChan {
			recv = append(recv, v)
			// Simulate processing the received value
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("receive channel routine end.")
	}()

	// Start sending data concurrently
	wg.Add(3) // We are launching 3 goroutines, so we need to wait for 3 tasks to complete
	go sendInt(mainChan, 10, &wg)
	go sendInt(mainChan, 100, &wg)
	go sendInt(mainChan, 1000, &wg)

	// Wait for all senders to complete and then close the channel
	go func() {
		wg.Wait()       // Wait for all sendInt goroutines to finish
		close(mainChan) // Close the channel when done
		fmt.Println("wait all send routine end. close mainChan.")

	}()

	// Wait for the receiver goroutine to finish (optional, depending on program logic)
	//wg.Add(1) // Adding one more task for receiving to wait for completion
	//go func() {
	//	// Just wait for all data to breceivere received
	//	fmt.Println("Additional wait routine started. wg.Done()")
	//	wg.Done() // Once the  goroutine finishes, signal that it's done
	//}()
	//wg.Wait() // Wait for the receiver to finish
	wgRecv.Wait() // wait or receiver to finish
	fmt.Printf("receive finished: recv = %v\n", recv)
}

func sendInt(c chan int, varStart int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it completes
	for i := 0; i < 10; i++ {
		c <- i + varStart
		fmt.Printf("send %d to channel\n", i+varStart)
	}
}

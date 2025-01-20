package main

import (
	"fmt"
	"sync"
)

func main() {
	// test server routine send data to one channel, main receive use slice
	mainChan := make(chan int, 4)
	var wg sync.WaitGroup

	// Start receiving data concurrently
	recv := make([]int, 0, 10) // Preallocate space for up to 30 elements
	go func() {
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

	/* chatgpt: 这部分代码的主要目的是确保主 goroutine 在程序结束前等待接收 goroutine 完成接收操作。我们可以逐行分析它的含义：
	我们启动了一个新的 goroutine，其中只有一条 wg.Done() 语句。这个 goroutine 的作用是：
	启动接收等待 goroutine：go func()
	完成任务：一旦接收任务（即接收 goroutine）完成，它通过 wg.Done() 来告知 sync.WaitGroup 这个任务已经结束了。
	在之前的代码中，接收操作是在一个独立的 goroutine 中进行的，它会从通道中接收数据，并在接收完所有数据后自动退出。然而，
	为了确保主程序等待这个接收操作的完成，我们显式地通过这个 go func() 来处理。
	实际上，这段代码并非每次都需要，具体是否加上它取决于你的应用需求。其目的是确保：

	接收操作的完成：主 goroutine 在输出最终结果之前，确保接收操作已经完成。
	同步收尾：如果接收操作是异步的，主程序可能会提前退出，导致输出未完成。因此，通过 sync.WaitGroup 来确保所有的接收操作完成，程序才退出。

	*/
	// Wait for the receiver goroutine to finish (optional, depending on program logic)
	wg.Add(1) // Adding one more task for receiving to wait for completion
	go func() {
		// Just wait for all data to breceivere received
		fmt.Println("Additional wait routine started. wg.Done()")
		wg.Done() // Once the  goroutine finishes, signal that it's done
	}()
	wg.Wait() // Wait for the receiver to finish
	fmt.Printf("receive finished: recv = %v\n", recv)
}

func sendInt(c chan int, varStart int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it completes
	for i := 0; i < 10; i++ {
		c <- i + varStart
		fmt.Printf("send %d to channel\n", i+varStart)
	}
}

/*
 为什么要再后面另外创建一个routine，去wg.Done,而不直接在接收routine中wg.Add(1),
 执行完for range接收所有数据后，接收routine wg.Done。 最后main中直接wg.wait呢?
 chatGPT:
你的理解是完全正确的，实际上你提到的方式更为简洁和自然。确实可以直接在接收 goroutine 中进行 wg.Add(1) 和 wg.Done() 操作，
而不需要额外再创建一个新的 goroutine 来调用 wg.Done()。

我原本加那个额外的 goroutine 主要是为了避免在接收 goroutine 中直接调用 wg.Add(1)，因为 wg.Add(1) 应该在主 goroutine
中执行，通常我们建议在启动 goroutine 之前就调用 Add，而不是在 goroutine 中调用。

但是在你的情境中，完全可以直接在接收 goroutine 中进行 wg.Add(1) 和 wg.Done()，并且这样实现更直观，代码更简洁。
*/
// 改进后的code：
//
//func main() {
//	// test server routine send data to one channel, main receive use slice
//	mainChan := make(chan int, 30)
//	var wg sync.WaitGroup
//
//	// Start receiving data concurrently
//	wg.Add(1) // Add 1 for the receiver goroutine
//	go func() {
//		defer wg.Done()            // Mark this goroutine as done when it completes
//		recv := make([]int, 0, 30) // Preallocate space for up to 30 elements
//		for v := range mainChan {
//			recv = append(recv, v)
//			// Simulate processing the received value
//			fmt.Printf("Received: %d\n", v)
//		}
//		fmt.Printf("receive finished: recv = %v\n", recv)
//	}()
//
//	// Start sending data concurrently
//	wg.Add(3) // We are launching 3 goroutines to send data, so we need to wait for 3 tasks to complete
//	go sendInt(mainChan, 10, &wg)
//	go sendInt(mainChan, 100, &wg)
//	go sendInt(mainChan, 1000, &wg)
//
//	// Wait for all senders to complete and then close the channel
//	go func() {
//		wg.Wait()       // Wait for all sendInt goroutines to finish
//		close(mainChan) // Close the channel when done
//	}()
//
//	// Wait for the receiver goroutine to finish (optional)
//	wg.Wait() // Wait for all goroutines (send and receive) to complete
//}

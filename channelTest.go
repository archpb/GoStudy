package main

import "fmt"

func main() {
	// 定义1：无缓冲channel var <channelName> chan <dataType>
	myNobufChan := make(chan int)
	fmt.Printf("myNobufChan len=%v, cap=%v\n", len(myNobufChan), cap(myNobufChan))
	fmt.Println("------------------------------------")

	// 定义2：有缓冲channel
	var myBufChan chan int = make(chan int, 8)

	fmt.Printf("myBufChan = %v(type:%T)\n", myBufChan, myBufChan)
	fmt.Printf("myBufChan len=%v, cap=%v\n", len(myBufChan), cap(myBufChan))
	myBufChan <- 1 // write data to channel
	fmt.Printf("send 1 int:  len=%v, cap=%v\n", len(myBufChan), cap(myBufChan))
	myBufChan <- 2
	fmt.Printf("send 2 int:  len=%v, cap=%v\n", len(myBufChan), cap(myBufChan))

	myBufChan <- 3
	fmt.Printf("send 3 int:  len=%v, cap=%v\n", len(myBufChan), cap(myBufChan))

	var receiver []int = make([]int, 8)
	fmt.Printf("type of receiver[0]:%T\n", receiver[0])
	receiver[0] = <-myBufChan
	fmt.Printf("receive 1 int:  len=%v, cap=%v\n", len(myBufChan), cap(myBufChan))

	receiver[1] = <-myBufChan
	receiver[2] = <-myBufChan
	fmt.Printf("receive 3 int:  len=%v, cap=%v\n", len(myBufChan), cap(myBufChan))

	fmt.Printf("receiver = %v\n", receiver)
	fmt.Println("------------------------------------")

	// define string channel
	var myStrChan chan string = make(chan string, 1024)
	fmt.Printf("myStrChan = %v(type:%T)\n", myStrChan, myStrChan)
	fmt.Printf("myStrChan len=%v, cap=%v\n", len(myStrChan), cap(myStrChan))
	myStrChan <- "aaa" // write data to channel
	fmt.Printf("send 1 int:  len=%v, cap=%v\n", len(myStrChan), cap(myStrChan))
	myStrChan <- "bbb"
	fmt.Printf("send 2 int:  len=%v, cap=%v\n", len(myStrChan), cap(myStrChan))

	myStrChan <- "ccc哈喽"
	fmt.Printf("send 3 int:  len=%v, cap=%v\n", len(myStrChan), cap(myStrChan))

	a := <-myStrChan
	b := <-myStrChan
	c := <-myStrChan

	fmt.Printf("a = %v, b = %v, c = %v\n", a, b, c)
	fmt.Println("------------------------------------")

	//test sever routine send data to one channel, main receive use slice
	mainChan := make(chan int, 20)
	go sendInt(mainChan, 10)
	go sendInt(mainChan, 100)
	go sendInt(mainChan, 1000)
	//read channel
	recv := make([]int, 40)
	for v := range mainChan {
		recv = append(recv, v)
	}
	fmt.Printf("receive finished: recv = %v\n", recv)
}

func sendInt(c chan int, varStart int) {
	for i := 0; i < 10; i++ {
		c <- i + varStart
		fmt.Printf("send %d to channel\n", i+varStart)
	}
	close
}

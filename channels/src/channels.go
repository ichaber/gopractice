package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	simpleChannel()
	bufferedChannel()
}

func simpleChannel() {
	channel := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Printf("The received number is %v\n", i)
		wg.Done()
	}(channel)
	go func(ch chan<- int) {
		ch <- 1000
		wg.Done()
	}(channel)
	wg.Wait()
}

func bufferedChannel() {
	channel := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		// Explicit for loop with check if channel is closed
		// for {
		// 	if i, ok := <-ch; ok {
		// 		fmt.Println(i)
		// 	} else {
		// 		break
		// 	}
		// }
		for i := range ch {
			fmt.Printf("The received number is %v\n", i)
		}
		wg.Done()
	}(channel)
	go func(ch chan<- int) {
		ch <- 111
		ch <- 999
		close(ch)
		wg.Done()
	}(channel)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This shows the channels demo")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()

	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		close(myCh)
		myCh <- 5
		wg.Done()

	}(myCh, wg)

	wg.Wait()

}

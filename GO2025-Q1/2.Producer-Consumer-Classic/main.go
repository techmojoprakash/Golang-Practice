package main

import (
	"fmt"
	"sync"
)

/*

Producer and Consumer Classic problem in Computer science

Producer will Produce the data to channel , consumer will consume from it.
this is not buffer channel. so it can keep 1 at a time, once it's consumed then Producer will keep another one.

Wait groups will help us to sync the program excecution here.

common channel is used to exchange the data between two go routines .

Note : Need to close the channel at the Producer end only. Channel should close otherwise deadlock.

after Producer and consumer func , it should do wg.Done .otherwise thread will not release....

here total 3 go routines running,
1. Main
2. Producer
3. Consumer

*/

var wg sync.WaitGroup

func Producer(centralChan chan int) {

	for i := 0; i < 10; i++ {
		centralChan <- i
		fmt.Printf("Producer no : %d\n", i)
	}
	close(centralChan)
	wg.Done()

}

func Consumer(centralChan chan int) {

	for i := range centralChan {
		fmt.Printf("Consumer no : %d\n", i)
	}
	wg.Done()

}

func main() {
	centralChan := make(chan int)
	fmt.Println("Program started...!")
	wg.Add(2)
	go Producer(centralChan)
	go Consumer(centralChan)
	wg.Wait()
	fmt.Println("DONE")

}

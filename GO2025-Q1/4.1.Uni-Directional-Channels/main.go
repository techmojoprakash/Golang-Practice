/*

First create channel when we need to write odd or even numbers to the channel by respective func's
Number center has to read the data from that channel

we should use Uni directional channels in all func's

*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func GenerateEventNum(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			ch <- i
			fmt.Println("Generated EventNum : ", i)
		}
	}
}

func GenerateOddNum(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			ch <- i
			fmt.Println("Generated OddNum : ", i)
		}
	}
}

func NumberCenter(ch <-chan int) {
	for i := range ch {
		fmt.Println("NumberCenter : ", i)
	}
}

func main() {
	fmt.Println("Uni Directional channel demo")

	public_chan := make(chan int)
	wg.Add(2)
	go GenerateOddNum(public_chan)
	go GenerateEventNum(public_chan)
	go func() {
		wg.Wait() // wait to finish GenerateOddNum and GenerateEventNum func,
		// then close channel
		close(public_chan)
	}()
	NumberCenter(public_chan)
}

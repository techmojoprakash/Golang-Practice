/*
Here we are simulating Multi channel communication

1. Medical and Engg students has independent channels to send marks to CentralBoard.
2. Central board will receive the data and print it.
Medi/Engg ==> CentralBoard

Many to One.... channel

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func getDelay() int {
	// Create a new source of randomness
	source_ra := rand.NewSource(time.Now().UnixNano())
	// Create a new random number generator
	r := rand.New(source_ra)

	return r.Intn(2) + 1

}

func sendMarks_Medical(medi_chan chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		delay := getDelay()
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Printf("Medical Marks sent : %d \n", i)
		medi_chan <- i
	}
	close(medi_chan)
}

func sendMarks_Engg(engg_chan chan int) {
	defer wg.Done()
	for i := 10; i < 15; i++ {
		delay := getDelay()
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Printf("Engg Marks sent : %d \n", i)
		engg_chan <- i
	}
	close(engg_chan)
}

// central board can receive the data from medical and enginnering marks and Print.
// in select, we should not return if any channel is closed. simply log and continue because other channels still open ...
func CentralBoard(medi_chan, engg_chan chan int) {
	fmt.Println("Central Board is start listening....!")
	defer wg.Done()

	isMediChanClosed := false
	isEnggChanClosed := false

	for {
		select {
		case medical_data, ok := <-medi_chan:
			if !ok {
				fmt.Println("Medi-Channel is closed")
				time.Sleep(time.Millisecond * 1000)
				isMediChanClosed = true
			} else {
				fmt.Printf("CentralBoard::Medi-Channel sent marks : %d\n", medical_data)

			}
		case engineer_data, ok := <-engg_chan:
			if !ok {
				fmt.Println("Enginnering-Channel is closed")
				time.Sleep(time.Millisecond * 1000)
				isEnggChanClosed = true
			} else {
				fmt.Printf("CentralBoard::Medi-Channel sent marks : %d\n", engineer_data)

			}

		default: // Handle the case where no channel is ready
			// You can add a timeout here using time.After
			// to avoid indefinite waiting if all channels are closed

			// Introduce a short delay to avoid excessive CPU usage
			time.Sleep(time.Millisecond * 1)
		}

		// Check if both channels are closed
		if isMediChanClosed && isEnggChanClosed {
			fmt.Println("All channels are closed. Exiting.")
			return
		}
		// Check if both channels are closed - NOT WORKING METHOD
		// if medi_chan ==nil && engg_chan == nil {
		// 	fmt.Println("All channels are closed. Exiting.")
		// 	return
		// }
	}
}
func main() {
	fmt.Println("Hello Multi Channel Communication...!")
	medi_chan := make(chan int)
	engg_chan := make(chan int)
	wg.Add(3)
	go sendMarks_Medical(medi_chan)
	go sendMarks_Engg(engg_chan)
	go CentralBoard(medi_chan, engg_chan) // CentralBoard will wait till all channels closed.
	wg.Wait()
	fmt.Println("Done with Main")
}

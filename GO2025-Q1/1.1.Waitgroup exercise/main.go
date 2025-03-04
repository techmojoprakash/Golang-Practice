package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
	defer wg.Done()
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()

	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()

	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()

	printMessage()

	fmt.Println("****DONE******")
}

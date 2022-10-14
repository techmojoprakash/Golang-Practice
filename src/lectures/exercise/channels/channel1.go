// Create Channel in Go

// channelName := make(chan int)

// channelName - name of the channel
// (chan int) - indicates that the channel is of integer type

// -----------1. Send data to the channel

// channelName <- data
// Here the data after the <- operator is sent to channelName.

// send integer data to channel
// number <- 15

// // send string data
// message <- "Learning Go Channel"

// ------------2. Receive data from the channel

// The syntax to receive data from the channel is:

// <- channelName
// This accesses the data from channelName.

// Let's see some example,

// // receive data 15
// <- number

// // receive data "Learning Go Channel"
// <- message
// -------------------------example-------------

// package main

// import "fmt"

// func main() {

// 	// create channel
// 	number := make(chan int)
// 	message := make(chan string)

// 	// function call with goroutine
// 	go channelData(number, message)

// 	// retrieve channel data
// 	fmt.Println("Channel Data:", <-number)
// 	fmt.Println("Channel Data:", <-message)

// }

// func channelData(number chan int, message chan string) {

// 	// send data into channel
// 	number <- 15
// 	message <- "Learning Go channel"

// }

// output

// Channel Data: 15
// Channel Data:  Learning Go Channel

// ---------------------example-2-------------
// Blocking Nature of Channel

// package main

// import "fmt"

// func main() {

// 	// create channel
// 	ch := make(chan string)

// 	// function call with goroutine
// 	go sendData(ch)

// 	// receive channel data
// 	fmt.Println(<-ch)

// }

// func sendData(ch chan string) {

// 	// data sent to the channel
// 	ch <- "Received. Send Operation Successful"
// 	fmt.Println("No receiver! Send Operation Blocked")

// }

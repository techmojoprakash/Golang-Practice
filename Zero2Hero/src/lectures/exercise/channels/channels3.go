// Deadlock

// One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

// Similarly, if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

// package main

// func main() {
//     ch := make(chan int)
//     ch <- 5
// }

// output:-

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//     /tmp/sandbox046150166/prog.go:6 +0x50

// ---------------------------second-----------------------------

// Unidirectional channels

// All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them. It is also possible to create unidirectional channels, that is channels that only send or receive data.

// package main

// import "fmt"

// func sendData(sendch chan<- int) {
//     sendch <- 10
// }

// func main() {
//     sendch := make(chan<- int)
//     go sendData(sendch)
//     fmt.Println(<-sendch)
// }

// output:-

// ./prog.go:12:14: invalid operation: <-sendch (receive from send-only type chan<- int)

// ----------------------------third------------------------------------------

// All is well but what is the point of writing to a send only channel if it cannot be read from!

// This is where channel conversion comes into use. It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.

// package main

// import "fmt"

// func sendData(sendch chan<- int) {
// 	sendch <- 10
// }

// func main() {
// 	chnl := make(chan int)
// 	go sendData(chnl)
// 	fmt.Println(<-chnl)
// }

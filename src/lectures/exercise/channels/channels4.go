// Closing channels and for range loops on channels

// Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.

// Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.

// v, ok := <- ch

// In the above statement ok is true if the value was received by a successful send operation to a channel. If ok is false it means that we are reading from a closed channel. The value read from a closed channel will be the zero value of the channel's type. For example, if the channel is an int channel, then the value received from a closed channel will be 0.

// package main

// import (
//     "fmt"
// )

// func producer(chnl chan int) {
//     for i := 0; i < 10; i++ {
//         chnl <- i
//     }
//     close(chnl)
// }
// func main() {
//     ch := make(chan int)
//     go producer(ch)
//     for {
//         v, ok := <-ch
//         if ok == false {
//             break
//         }
//         fmt.Println("Received ", v, ok)
//     }
// }

// -----------------second- string----------------
// The for range form of the for loop can be used to receive values from a channel until it is closed.

// package main

// import (
// 	"fmt"
// )

// func producer(chnl chan int) {
//     for i := 0; i < 10; i++ {
//         chnl <- i
//     }
//     close(chnl)
// }
// func main() {
//     ch := make(chan int)
//     go producer(ch)
//     for v := range ch {
//         fmt.Println("Received ",v)
//     }
// }

// -----------------------------third-------------------

package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}
}

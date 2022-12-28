// sample program with channels and goroutines
// package main

// import (
// 	"fmt"
// )

// func hello(done chan bool) {
// 	fmt.Println("Hello world goroutine")
// 	done <- true
// }
// func main() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<-done
// 	fmt.Println("main function")
// }

// output

// Hello world goroutine
// main function

// -----------------------------------second------------------------------

// // program will print the sum of the squares and cubes of the individual digits of a number.

// // For example, if 123 is the input, then this program will calculate the output as

// // squares = (1 * 1) + (2 * 2) + (3 * 3)
// // cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
// // output = squares + cubes = 50
// // We will structure the program such that the squares are calculated in a separate Goroutine, cubes in another Goroutine and the final summation happens in the main Goroutine.

package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

// // output

// Final output 1536

// ------------------------third-------------------------

// package main

// import (
// 	"fmt"
// )

// func digits(number int, dchnl chan int) {
// 	for number != 0 {
// 		digit := number % 10
// 		dchnl <- digit
// 		number /= 10
// 	}
// 	close(dchnl)
// }
// func calcSquares(number int, squareop chan int) {
// 	sum := 0
// 	dch := make(chan int)
// 	go digits(number, dch)
// 	for digit := range dch {
// 		sum += digit * digit
// 	}
// 	squareop <- sum
// }

// func calcCubes(number int, cubeop chan int) {
// 	sum := 0
// 	dch := make(chan int)
// 	go digits(number, dch)
// 	for digit := range dch {
// 		sum += digit * digit * digit
// 	}
// 	cubeop <- sum
// }

// func main() {
// 	number := 589
// 	sqrch := make(chan int)
// 	cubech := make(chan int)
// 	go calcSquares(number, sqrch)
// 	go calcCubes(number, cubech)
// 	squares, cubes := <-sqrch, <-cubech
// 	fmt.Println("Final output", squares+cubes)
// }

// output:-
// Final output 1536

// //--Summary:
// //  Implement receiver functions to create stat modifications
// //  for a video game character.
// //
// //--Requirements:
// //* Implement a player having the following statistics:
// //  - Health, Max Health
// //  - Energy, Max Energy
// //  - Name
// //* Implement receiver functions to modify the `Health` and `Energy`
// //  statistics of the player.
// //  - Print out the statistic change within each function
// //  - Execute each function at least once

// package main

// import "fmt"

// type MySqure struct {
// 	length int
// }

// // receiver function - method - 1
// func (sq *MySqure) Area() {
// 	sq.length = sq.length * 2
// }

// // receiver function - method - 2
// func Area(sq *MySqure) {
// 	sq.length = sq.length * 2
// }

// func main() {
// 	fmt.Println("Hello Receiver functions concept")
// 	var hydSq MySqure
// 	hydSq.length = 2
// 	// receiver function - method - 1
// 	// function calling
// 	// hydSq.Area()

// 	// receiver function - method - 2
// 	// function calling
// 	Area(&hydSq)
// 	fmt.Println("New Area : ", hydSq.length)

// }

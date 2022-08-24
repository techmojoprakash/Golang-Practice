package main

import "fmt"

// if we mention defer then , just that statement/block will skip and excecute at end of the program
// usually it was stored in stack - LIFO rule
// defer can also apply for function calls too

// func main() {
// 	fmt.Println("Hello") // 1
// 	defer fmt.Println("Techmojo") // 4
// 	fmt.Println("Hyderabad") // 2
// 	defer fmt.Println("This is ") //3

// }

// --------------------------------------

// placing defer with in loop
// func PrintData() {
// 	i := 0
// 	for i < 5 {
// 		defer fmt.Println(i) // - here stack 01234, then print 43210
// 		i++
// 	}
// }

// func main() {
// 	fmt.Println("Hello")          // 1
// 	defer fmt.Println("Techmojo") // 5
// 	PrintData()                   // 4 - here stack 01234, then print 43210
// 	fmt.Println("Hyderabad")      // 2
// 	defer fmt.Println("This is ") //3

// }

// --------------------------------------

// placing defer outside of loop
func PrintOutsideData() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func main() {
	fmt.Println("Hello")          // 1
	defer fmt.Println("Techmojo") // 5
	defer PrintOutsideData()      // 4 - no stack here just excecute this
	fmt.Println("Hyderabad")      // 2
	defer fmt.Println("This is ") //3

}

// defer main useful case
// func main(){
// 	file.open() // open file
// 	defer file.close() // close file with defer
// 	// means after all statements excecute above defer to close file
// 	file.print()
// }

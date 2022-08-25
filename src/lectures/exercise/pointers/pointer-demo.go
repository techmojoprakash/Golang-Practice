// package main

// import "fmt"

// func increment(xPointer *int) {
// 	// accepting pointer and dereferenced and increment value by 1
// 	// simply indirectly updating the actual value based on it's address
// 	// *xPointer += 1
// 	*xPointer = *xPointer + 1
// }

// func main() {
// 	fmt.Println("Hello")
// 	x := 123
// 	// method - 1
// 	increment(&x)
// 	fmt.Println(x)

// 	// // method - 2
// 	// var xPointer *int   // Pointer declaration
// 	// xPointer = &x       // pointer address assignment
// 	// increment(xPointer) // passing pointer to method
// 	// fmt.Println(x)

// }

// // Value receivers in methods vs Value arguments in functions

// // When a function has a value argument, it will accept only a value argument.

// // When a method has a value receiver, it will accept both pointer and value receivers.

// // Let's understand this by means of an example.

// package main

// import (
// 	"fmt"
// )

// type rectangle struct {
// 	length int
// 	width  int
// }

// func area(r rectangle) {
// 	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
// }

// func (r rectangle) area() {
// 	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
// }

// func main() {
// 	r := rectangle{
// 		length: 10,
// 		width:  5,
// 	}
// 	area(r)
// 	r.area()

// 	p := &r
// 	/*
// 	   compilation error, cannot use p (type *rectangle) as type rectangle
// 	   in argument to area
// 	*/
// 	//area(p)

// 	p.area() //calling value receiver with a pointer
// }

// //prakash notes: method can accept both value and reference but incase of func only value.

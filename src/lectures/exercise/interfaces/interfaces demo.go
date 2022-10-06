// package main

// import "fmt"

// //  ***NOTE***
// // you can pass values two ways
// // 1. pass by value 2.Pass by reference
// // don't mix above two in one interface

// // creating an interface for shapes rectangle, square
// type AreaAndPeremeter interface {
// 	Area() float64
// 	Peremeter() float64
// }

// // structure for rectangle
// type rectangle struct {
// 	width  float64
// 	height float64
// }

// // structure for squre
// type squre struct {
// 	side float64
// }

// // methods for rectangle
// func (r rectangle) Peremeter() float64 {
// 	return 2 * (r.height + r.width)
// }

// func (r rectangle) Area() float64 {
// 	return r.height * r.width
// }

// // methods for square
// func (s squre) Peremeter() float64 {
// 	return s.side * 4
// }

// func (s squre) Area() float64 {
// 	return s.side * s.side
// }

// func main() {
// 	// creating var for a struct
// 	rect := rectangle{width: 5, height: 4}
// 	sq := squre{side: 3}

// 	// calling Area and Rectangle methods
// 	fmt.Println(" Rectangle Area: ", rect.Area())
// 	fmt.Println(" Rectangle Peremeter: ", rect.Peremeter())

// 	fmt.Println(" Squre Area: ", sq.Area())
// 	fmt.Println(" Squre Peremeter: ", sq.Peremeter())

// 	// var i string = "12"
// 	// type j string
// 	// var ss2 string
// 	// var ss j = "hello"
// 	// fmt.Println("ss", ss)

// }

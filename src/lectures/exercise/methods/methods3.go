// // Methods of anonymous struct fields

// // Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where the anonymous field is defined.

// package main

// import (
// 	"fmt"
// )

// type address struct {
// 	city  string
// 	state string
// }

// func (a address) fullAddress() {
// 	fmt.Printf("Full address: %s, %s", a.city, a.state)
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	address
// }

// func main() {
// 	p := person{
// 		firstName: "Elon",
// 		lastName:  "Musk",
// 		address: address{
// 			city:  "Los Angeles",
// 			state: "California",
// 		},
// 	}

// 	p.fullAddress() //accessing fullAddress method of address struct

// }

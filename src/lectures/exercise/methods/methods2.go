// Pointer Receivers vs Value Receivers

// package main

// import (
// 	"fmt"
// )

// type Employee struct {
// 	name string
// 	age  int
// }

// /*
// Method with value receiver
// */
// func (e Employee) changeName(newName string) {
// 	e.name = newName
// }

// /*
// Method with pointer receiver
// */
// func (e *Employee) changeAge(newAge int) {
// 	e.age = newAge
// }

// func main() {
// 	e := Employee{
// 		name: "Mark Andrew",
// 		age:  50,
// 	}
// 	fmt.Printf("Employee name before change: %s", e.name)
// 	e.changeName("Michael Andrew")
// 	fmt.Printf("\nEmployee name after change: %s", e.name)

// 	fmt.Printf("\n\nEmployee age before change: %d", e.age)
// 	(&e).changeAge(51)
// 	fmt.Printf("\nEmployee age after change: %d", e.age)
// }

// ------output-------
// Employee name before change: Mark Andrew
// Employee name after change: Mark Andrew

// Employee age before change: 50
// Employee age after change: 51

// When to use pointer receiver and when to use value receiver:-

// Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.

// Pointers receivers can also be used in places where it's expensive to copy a data structure. Consider a struct that has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied which will be expensive. In this case, if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.
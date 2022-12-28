package main

import "fmt"

func main() {
	fmt.Println("Just started excecution of program")
	// creating variable
	dell := 231
	// creating pointer for dell
	dellPointer := &dell

	// we need to put * to get actual value of Pointer
	fmt.Println("dellPointer: ", dellPointer)                   //geting address
	fmt.Println("actual value(dereferencing) : ", *dellPointer) //getting value in that address

	// change value based on pointer
	*dellPointer = 567
	fmt.Println("dell new value: ", dell)

}

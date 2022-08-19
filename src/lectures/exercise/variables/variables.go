//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var myFavColor = "Red"
	fmt.Println(myFavColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	year, age := 2019, 28
	fmt.Println(year, age)

	//* Store your first & last initials in two variables using block assignment

	var (
		firstName = "Jaya"
		lastName  = "Prakash"
	)
	fmt.Println(firstName, lastName)

	newName, lastName := "Aluri", "JAYA"

	fmt.Println(newName, lastName)

}

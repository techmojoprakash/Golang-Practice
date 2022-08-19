//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func putName(myName string) string {
	result := "Your Name is " + myName
	return result
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

func mySum(fir, sec, thi int) int {
	return fir + sec + thi
}

// * Write a function that returns any number
func getNum(dellNum int) int {
	return dellNum
}

//* Write a function that returns any two numbers

func getTwoNum(firstNo, secondNo int) (int, int) {
	return firstNo, secondNo
}

// * Add three numbers together using any combination of the existing functions.
var result = mySum(1, 2, 3) + getNum(4)

//  * Print the result
//* Call every function at least once

func main() {
	fmt.Println(putName("Goa"))
	fmt.Println("My Sum is", mySum(1, 2, 3))
	fmt.Println("My Num is", getNum(123))
	var ak, pk = getTwoNum(12, 3) // two values return
	fmt.Println("Return 2 Num is", ak, pk)
	fmt.Println("result", result)

}

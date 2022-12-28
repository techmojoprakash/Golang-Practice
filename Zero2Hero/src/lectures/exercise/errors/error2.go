// // 2. Error using Errorf() in Golang

// // We can also handle Go errors using the Errorf() function. Unlike, New(), we can format the error message using Errorf().

// // This function is present inside the fmt package, so we can directly use this if we have imported the fmt package.

// package main

// import "fmt"

// func main() {

// 	age := -14

// 	// create an error using Efforf()
// 	error := fmt.Errorf("%d is negative\nAge can't be negative", age)

// 	if age < 0 {
// 		fmt.Println(error)
// 	} else {
// 		fmt.Println("Age: %d", age)
// 	}
// }

// // --------------output----------
// // OUTPUT

// // -14 is negative
// // Age can't be negative

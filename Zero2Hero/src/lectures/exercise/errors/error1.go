// // 1. Go Error using New() Function

// // In Go, we can use the New() function to handle an error.
// // This function is defined inside the errors package and allows us to create our own error message.

// package main

// // import the errors package
// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	message := "Hello"

// 	// create error using New() function
// 	myError := errors.New("WRONG MESSAGE")

// 	if message != "Programiz" {
// 		fmt.Println(myError)
// 	}

// }

// // OUTPUT
// // WRONG MESSAGE

// // ---------------------------------SECOND one---------------------------------------------

// package main

// // import the errors package
// import (
// 	"errors"
// 	"fmt"
// )

// // function that checks if name is Programiz
// func checkName(name string) error {

// 	// create a new error
// 	newError := errors.New("Invalid Name")

// 	// return the error if name is not Programiz
// 	if name != "Programiz" {
// 		return newError
// 	}

// 	// return nil if there is no error
// 	return nil
// }

// func main() {

// 	name := "Hello"

// 	// call the function
// 	err := checkName(name)

// 	// check if the err is nil or not
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Valid Name")
// 	}

// }

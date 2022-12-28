/* Def : Label is used in break and continue statement where it’s optional but it’s required in goto statement. Scope of the label is the function where it’s declared. It doesn’t start after the declaration though as with (short) variable declaration but is available for the whole function body:
 */

// Syntax------->

// label:
//// Code to execute in label

// ----------------->

// -----------------example-1 ------------------

// note : lable should be with in function but if we mention outside of the function we will get error : "label End not defined"
// if it's unused then , error : “label End defined and not used”

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello, Reader! Your learning about 'goto' statement")
// 	// We create a for loop which runs until i is 10
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("Index: %d\n", i)
// 		if i == 5 {
// 			// When i is 5, lets exit by using goto
// 			goto exit
// 		}
// 	}
// 	fmt.Println("Skip this line here")
// 	// Create the exit label and insert code that should be executed when triggered
// exit:
// 	fmt.Println("We are now exiting the program")
// }

// -----------------example-2 ------------------

// in the below code we are calling exit from out side of declared function, then you will get error : "label exit not defined"

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Reader! Your learning about 'goto' statement")
	// We create a for loop which runs until i is 10
	for i := 0; i < 10; i++ {
		// printInteger will try to goto to exit, which will not work
		printInteger(i)
	}
	fmt.Println("Skip this line here")
	// Create the exit label and insert code that should be executed when triggered
exit:
	fmt.Println("We are now exiting the program")
}

func printInteger(i int) {
	fmt.Printf("Index: %d\n", i)
	if i == 5 {
		// When i is 5, lets exit by using goto
		goto exit
	}
}

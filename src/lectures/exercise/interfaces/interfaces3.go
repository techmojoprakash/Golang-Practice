// Type assertion

// Type assertion is used to extract the underlying value of the interface.

// i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.

// package main

// import (
//     "fmt"
// )

// func assert(i interface{}) {
//     v, ok := i.(int)
//     fmt.Println(v, ok)
// }
// func main() {
//     var s interface{} = 56
//     assert(s)
//     var i interface{} = "Steven Paul"
//     assert(i)
// }

// output:-
// 56 true
// 0 false

// --------------Type switch--------------------

// package main

// import (
// 	"fmt"
// )

// func findType(i interface{}) {
// 	switch i.(type) {
// 	case string:
// 		fmt.Printf("I am a string and my value is %s\n", i.(string))
// 	case int:
// 		fmt.Printf("I am an int and my value is %d\n", i.(int))
// 	default:
// 		fmt.Printf("Unknown type\n")
// 	}
// }
// func main() {
// 	findType("Naveen")
// 	findType(77)
// 	findType(89.98)
// }

// output:-

// I am a string and my value is Naveen
// I am an int and my value is 77
// Unknown type  
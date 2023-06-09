package main

import "fmt"

func main() {

	fmt.Println("Hello")
}

// Creation of empty map in Golang
// 2 methods

// users : map[string]int{}
// users : make(map[string]int)

// The first statement users := map[string]int{} uses a shorthand syntax to create an empty map and assigns it to the variable users. This syntax is commonly used when you want to declare and initialize a map in a single line.
// The second statement users := make(map[string]int) uses the make function to create a map and assigns it to the variable users. The make function is a built-in function in Go that is used to create slices, maps, and channels. It takes the type of the map (map[string]int in this case) as the argument and returns an initialized and ready-to-use map.
// difference :-
// In terms of functionality, both approaches result in the creation of an empty map. However, the make function allows you to specify additional parameters, such as the initial capacity of the map, whereas the shorthand syntax does not provide that flexibility. If you don't specify a capacity, both methods create a map with an initial capacity of 0, which will dynamically grow as you add elements to it.

// Example

// package main

// import "fmt"

// func main() {
// 	// Using shorthand syntax
// 	users := map[string]int{}
// 	users["John"] = 25
// 	users["Alice"] = 30

// 	fmt.Println(users) // Output: map[Alice:30 John:25]

// 	// Using make function
// 	books := make(map[string]int)
// 	books["Harry Potter"] = 5
// 	books["The Great Gatsby"] = 3

// 	fmt.Println(books) // Output: map[Harry Potter:5 The Great Gatsby:3]

// 	// Using make function with initial capacity
// 	scores := make(map[string]int, 100)
//     	scores["Alice"] = 85
//     	scores["Bob"] = 92

//     fmt.Println(scores) // Output: map[Alice:85 Bob:92]
// }

// Example 2: Creating a map of maps:
// package main

// import "fmt"

// func main() {
//     records := make(map[string]map[string]string)

//     student1 := make(map[string]string)
//     student1["name"] = "Alice"
//     student1["major"] = "Computer Science"

//     student2 := make(map[string]string)
//     student2["name"] = "Bob"
//     student2["major"] = "Mathematics"

//     records["001"] = student1
//     records["002"] = student2

//     fmt.Println(records) // Output: map[001:map[major:Computer Science name:Alice] 002:map[major:Mathematics name:Bob]]
// }

// Example 3: Creating a map of slices:

// package main

// import "fmt"

// func main() {
//     groups := make(map[string][]string)

//     groups["team1"] = []string{"Alice", "Bob", "Charlie"}
//     groups["team2"] = []string{"Dave", "Eve"}

//     fmt.Println(groups) // Output: map[team1:[Alice Bob Charlie] team2:[Dave Eve]]
// }

// Video- 3

// Variables
// 	global variables
// 	local variables
// Constants
// Variable declarations

// Global Variable : if variable declare variable out side of function, we can say that it's global variable

// declare and assign
// here it will infer the type - it only works in global space
// var nameX = "Foo"
// var cityX = "bar"

// // we can explicitly specify types too
// var nameX string = "Foo"

// // we can group above ones
// var (
//     nameX = "Foo"
//     cityX = "bar"
// )

// // just delcare variable & assign default value ie: ""
// var nameX string

// Local variables

// // local

// func main(){
//     version := 2 // infer int
//     // it will infer datatype automatically

//     // alternative
//     var version = 2

//     // we can also declare inside with default value ie : 0
//     var version int

// }

// Constants

// // delcare
// const version = 1
// // const is always immutable

// constant need not be in upper case.

// Video-4 - Built-in & custom types

// // struct create
// type Player struct {
//     name string
//     health int
//     attackPower float64
// }

// //create object with default values

// func main() {
//     // player1 obj will create with default values
//     player1 := Player{}
//     // player2 obj will create with assigned values & attackPower will take default value
//     player2 := Player {
//         name : "Captain Jack",
//         health : 45
//         }

// }

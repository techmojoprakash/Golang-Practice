// //--Summary:
// //  Create a program to manage parts on an assembly line.
// //
// //--Requirements:
// //* Using a slice, create an assembly line that contains type Part
// //* Create a function to print out the contents of the assembly line
// //* Perform the following:
// //  - Create an assembly line having any three parts
// //  - Add two new parts to the line
// //  - Slice the assembly line so it contains only the two new parts
// //  - Print out the contents of the assembly line at each step
// //--Notes:
// //* Your program output should list 3 parts, then 5 parts, then 2 parts

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// // type Part string

// func main() {
// 	MySlice := []string{"tom", "cat", "rat", "fat", "sweet"}
// 	MySlice = append(MySlice, "meow")
// 	fmt.Println(MySlice)
// 	fmt.Println(MySlice[2])

// 	// sort slice
// 	sort.Strings(MySlice)
// 	fmt.Println("MySlice ", MySlice)

// 	// check sorted ?
// 	fmt.Println("is sorted := ", sort.StringsAreSorted(MySlice))

// }

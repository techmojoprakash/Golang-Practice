// Empty interface

// An interface that has zero methods is called an empty interface. It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface.

// package main

// import (
// 	"fmt"
// )

// func describe(i interface{}) {
// 	fmt.Printf("Type = %T, value = %v\n", i, i)
// }

// func main() {
// 	s := "Hello World"
// 	describe(s)
// 	i := 55
// 	describe(i)
// 	strt := struct {
// 		name string
// 	}{
// 		name: "Naveen R",
// 	}
// 	describe(strt)
// }

// Prakash Note : we can pass any type as parameters
// output:-

// Type = string, value = Hello World
// Type = int, value = 55
// Type = struct { name string }, value = {Naveen R}


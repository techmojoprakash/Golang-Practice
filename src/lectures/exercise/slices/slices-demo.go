// package main

// import "fmt"

// func main() {
// 	// creating slice
// 	// if we didn't mention array size, then it's simply called as slice.
// 	// here slices are similar to Arraylist in Java. ie. dynamic allocation of memory

// 	MySlice := []string{"tom", "cat", "rat", "fat", "sweety"}
// 	// append is function we need to pass two arguments 1.slice, 2.string that you want to add
// 	MySlice = append(MySlice, "meow", "livpure")
// 	fmt.Println("new slice is : ", MySlice)
// 	// index the slice
// 	fmt.Println(MySlice[2])
// 	// slice the slice
// 	fmt.Println(MySlice[2:4])

// 	// remove "fat" word from slice
// 	// so first slice the data from existing slice and remove the value
// 	// we need to place ... to accept multi arguments
// 	newSlice := append(MySlice[:2], MySlice[3:]...)
// 	fmt.Println("newSlice without rat => ", newSlice)

// 	// make slice and assign value later
// 	// in below we are giving the size for the slice like array
// 	// but we can call this as slice only
// 	// if we want to add more values to this slice, simply use append function.
// 	// append function will recrate the slice and add more elements to it.
// 	// Purpose : Improve the perfomence by giving fixed size initially

// 	freshSlice := make([]string, 4)
// 	freshSlice[0] = "first"
// 	freshSlice[1] = "second"
// 	freshSlice[2] = "third"
// 	freshSlice[3] = "fourth"

// 	fmt.Println("freshSlice is : ", freshSlice)
// 	freshSlice = append(freshSlice, "fifth", "six", "seventh")
// 	fmt.Println("Modified freshSlice is : ", freshSlice)

// 	//concatenate two slices using ... 3dot notation
// 	resultSlice := append(newSlice, freshSlice...)
// 	fmt.Println("Modified resultSlice is : ", resultSlice)

// }

package main

import "fmt"

// range can work like for each loop
// you can use this on array, slice, map, string

func main() {
	fmt.Println("Hello Range Example")

	// range on arryas
	myArray := [3]string{"a", "b", "c"}
	for i, j := range myArray {
		fmt.Printf("index is %v and value is %v \n", i, j)
	}

	// range on slices
	mySlice := []string{"a", "b", "c"}
	for i, j := range mySlice {
		fmt.Printf("index is %v and value is %v \n", i, j)
	}

	// range on maps
	myMap := map[int]string{1: "tcs", 2: "wipro", 3: "apple"}
	for i, j := range myMap {
		fmt.Printf("key is %v and value is %v \n", i, j)
	}

	// range on string
	myString := "Techmojo"
	for i, j := range myString {
		// fmt.Printf("index is %v and value is %v \n", i, j) // we can get ascii values
		// so we can use runs %q
		// %q -	a single-quoted character literal safely escaped with Go syntax.
		fmt.Printf("index is %v and value is %q \n", i, j) // we can get ascii values

	}

}

package main

import "fmt"

// meow := "Sweet"
// fmt.Println(meow)

func getTenNums() {
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8} // remaining values assigned to zeros
	for i := 0; i < len(myArray); i++ {
		fmt.Print(myArray[i])
	}
	fmt.Println()
}

func main() {
	getTenNums()
	// array declaration
	var firstArray [3]int
	firstArray[0] = 12
	firstArray[1] = 13
	firstArray[2] = 14

	// array declaration with values
	secondArray := [3]int{99, 88, 77}

	// array declaration with default values
	secondNewArray := [5]int{919, 883, 747}

	// array declaration with variable inputs
	thirdArray := [...]int{12, 13, 14, 15}

	//printing arrays
	fmt.Println(firstArray)

	fmt.Println(secondArray)

	fmt.Println(secondNewArray)

	fmt.Println(thirdArray)

}

package main

import "fmt"

func main() {
	fmt.Println("Hello")
	myArray := [5]int{6, 7, 8, 9, 10}
	for i := 10; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}

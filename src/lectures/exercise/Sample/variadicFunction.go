package main

import "fmt"

// if you don't know how many params , use this

func getSum(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println("Hello Variadics")
	fmt.Println("sum is : ", getSum(1, 2, 3, 4, 5))

	first := []int{1, 2, 3, 4, 5}
	second := []int{1, 2, 3, 4, 5}
	third := append(first, second...)
	fmt.Println("third", third)
	// we should not pass third slice directly
	// so place ... after varibale then it will unpack and send to func

	fmt.Println("getSum(third...) ", getSum(third...))

	// we could not send below way
	// fmt.Println("another way ", getSum(first, second...))

}

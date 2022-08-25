package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// variadic function
// Select single argument from all arguments of variadic function.
func varadicFunc(s ...string) {
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])

}

// variable args with varadic function
// // we can only use one variable inputs(ie => ...) in func & that should be last argument
// not working
// func varArgsSum(mystr string, myint ...int) int {
// 	i := 0
// 	sum := 0

// 	for myint[i] != 0 {
// 		sum += myint[i]
// 	}
// 	fmt.Println("varArgsSum printing", sum)
// 	return sum
// }

func main() {
	fmt.Println("Hello")

	// Declaring (Creating) Constants
	const PRODUCT string = "Canada"
	const PRICE = 500
	// Multilple Constants Declaration Block
	const (
		PRODUCT1 = "Mobile"
		QUANTITY = 50
		PRICE1   = 50.50
		STOCK    = true
	)
	fmt.Println(PRODUCT, PRICE, PRODUCT1, QUANTITY, PRICE1, STOCK)

	// Conversion - Type casting

	// How to Convert string to integer type in Go?
	// Convert string method -1
	// You can use the strconv packages Atoi() function to convert the string into an integer value. Atoi stands for ASCII to integer. The Atoi() function returns two values: the result of the conversion, and the error (if any).

	// Syntax
	// func Atoi(s string) (int, error)

	strVal := "100"
	intVar, err := strconv.Atoi(strVal)
	fmt.Println(" int data of string '100' is ", intVar, reflect.TypeOf(intVar), "error is ", err)
	// Convert string method -2

	// ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i. This function accepts a string parameter, convert it into a corresponding int type based on a base parameter. By default, it returns Int64 value.

	// Syntax
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)

	myIntvar, errval := strconv.ParseInt(strVal, 0, 32)
	fmt.Println(" ParseInt data of string '100' is ", myIntvar, reflect.TypeOf(myIntvar), "error is ", errval)

	// variadic function call
	varadicFunc("red", "green", "blue", "white")
	// variable args with varadic function
	// fmt.Println(varArgsSum("Rect", 123, 876))
}

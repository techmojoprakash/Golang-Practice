//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import (
	"fmt"
	"time"
)

// func main() {
// 	switch age := 33; {
// 	case age == 0:
// 		fmt.Println("newborn")

// 	case 0 <= age && 3 >= age:
// 		fmt.Println("toddler")

// 	case 4 <= age && 12 >= age:
// 		fmt.Println("child")

// 	case 13 <= age && 17 >= age:
// 		fmt.Println("teenager")

// 	default:
// 		fmt.Println("adult")
// 	}
// }

func main() {
	weekday := (time.Now().Weekday().String())
	// above case we need to use string otherwise we will get time format only
	// fmt.Println(reflect.TypeOf(weekday))

	// 1. Go switch case with fallthrough

	// Normal
	switch weekday == "Wednesday" {
	case true:
		fmt.Println("Yahooo! it's Wednesday")
	default:
		fmt.Println("Sorry it's not a Wednesday")

	}

	// 2. Go switch with multiple cases
	// multiple cases
	switch weekday {
	case "Monday", "Wednesday":
		fmt.Println("Yahooo! it's Wednesday or Monday")
	case "Tuesday", "Thursday":
		fmt.Println("Yahooo! it's Odd Day")
	case "Friday":
		fmt.Println("Yahooo! it's Last Working day of the week")
	default:
		fmt.Println("Sorry it's Weekend")

	}

	// 3. Golang switch without expression
	// In Go, the expression in switch is optional.
	// If we don't use the expression, the switch statement
	// is true by default.

	noOfDays := 28
	switch {
	case noOfDays == 28:
		fmt.Println("It's non Leap year")
	default:
		fmt.Println("It's Leap year")

	}

	// 4. Go switch optional statement
	// In Golang, we can also use an optional statement along with the expression. The statement and expression are separated by semicolons.

	switch myday := 4; myday {
	case 1:
		fmt.Println("day 1")
	case 2:
		fmt.Println("day 2")
	case 3:
		fmt.Println("day 3")
	default:
		fmt.Println("day 4")
	}

}

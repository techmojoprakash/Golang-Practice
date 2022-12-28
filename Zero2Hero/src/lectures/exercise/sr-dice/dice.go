//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func getEachInstance(totalDice, sidesVal int) int {
	sumVal := 0
	for {
		if totalDice <= 0 {
			break
		}
		dieVal := rand.Intn(sidesVal-1) + 1
		// dieVal = rand.Intn(sidesVal-1) + 1
		// dieVal = rand.Intn(sidesVal-1) + 1
		// dieVal = rand.Intn(sidesVal-1) + 1
		sumVal += dieVal
		totalDice--
	}
	return sumVal
}

func getnoOfTimestoRoll(noOfTimestoRoll, totalDice, sidesVal int) int {
	totalSum := 0
	for i := 0; i < noOfTimestoRoll; i++ {
		eachDie := getEachInstance(totalDice, sidesVal)
		totalSum += eachDie
	}
	return totalSum
}

func main() {
	noOfTimestoRoll, totalDice, sidesVal := 1, 1, 6
	totalRoll := getnoOfTimestoRoll(noOfTimestoRoll, totalDice, sidesVal)

	if totalRoll == 2 && totalDice == 2 {
		fmt.Println("Snake eyes : when the total roll is", totalRoll, "and total dice is", totalDice)
	}
	if totalRoll == 7 {
		fmt.Println("Lucky 7 when the total roll is", totalRoll)

	}
	if totalRoll%2 == 0 {
		fmt.Println("Even when the total roll is even", totalRoll)

	}

	if totalRoll%2 == 1 {
		fmt.Println("Odd when the total roll is odd", totalRoll)

	}

	// syntax
	// rand.Intn(max-min) + min
	// fmt.Println(rand.Intn(6-1) + 1)

}

// // First Go program
// package main

// import "fmt"

// func myCalc(firstval, secondval int) int {
// 	return firstval + secondval
// }

// const MyTmNumber = 12345

// // Main function
// func main() {
// 	// var myName = "Google"           // variable assign
// 	// var myNameis string = "Yahooo!" // variable assign
// 	// newName := "Google GO language" // variable assign
// 	// fmt.Println("!... Hello World ...!", newName, myName, myNameis)
// 	// fmt.Println("Hello", newName)

// 	// var mySum int // un-assigned variable
// 	// fmt.Println("mySum", mySum)

// 	// first, second := 1, 2 // simple way to create variables
// 	// fmt.Println(first, second)

// 	// // ok error idiom
// 	// third, second := 1, 4 // second created 2 times , still valid
// 	// // if you assign second saparately without another varibale , then it throws error
// 	// fmt.Println(third, second)

// 	// mySum = first + second + third
// 	// fmt.Println("mySum is", mySum)

// 	// // assign multiple variables in one go
// 	// var (
// 	// 	lessonName = "Science"
// 	// 	lessonDate = "Thursday"
// 	// )
// 	// fmt.Println(lessonName, lessonDate)

// 	// // ignoring variables
// 	// word1, word2, _ := "TM", "HYD", "117"
// 	// fmt.Println(word1, word2)

// 	// result := myCalc(1, 3)
// 	// fmt.Println(result, "result")

// 	// Random num generate
// 	// rand.Seed(time.Now().UnixNano())
// 	// myval := rand.Intn(1033)
// 	// fmt.Println(myval)
// 	// myval = rand.Intn(10)
// 	// fmt.Println(myval)

// 	// Type of
// 	// mydata := 2
// 	// fmt.Println(reflect.TypeOf(mydata))

// 	// type conversion
// 	// floatData := float32(mydata)
// 	// fmt.Println(reflect.TypeOf(floatData), floatData)

// 	// constants
// 	// Constants cannot be declared using the := syntax.
// 	fmt.Println(MyTmNumber)
// 	const MyTmNumber = 123452
// 	fmt.Println(MyTmNumber)
// }

package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// func needInt(x int) int { return x*10 + 1 }
// func needFloat(x float64) float64 {
//	  return x * 0.1

func main() {
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
	fmt.Println(int(Big))
	fmt.Println(Small)
}

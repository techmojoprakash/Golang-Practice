// // --Summary:
// //  Create a program that can activate and deactivate security tags
// //  on products.

// // --Requirements:
// // * Create a structure to store items and their security tag state - done
// //  - Security tags have two states: active (true) and inactive (false) - done
// // * Create functions to activate and deactivate security tags using pointers
// // * Create a checkout() function which can deactivate all tags in a slice
// // * Perform the following:
// //  - Create at least 4 items, all with active security tags
// //  - Store them in a slice or array
// //  - Deactivate any one security tag in the array/slice
// //  - Call the checkout() function to deactivate all tags
// //  - Print out the array/slice after each change

// package main

// import (
// 	"fmt"
// )

// // * Create a structure to store items and their security tag state - done
// //   - Security tags have two states: active (true) and inactive (false) - done
// type SecurityTags struct {
// 	item      string
// 	tagStatus bool
// }

// // * Create functions to activate and deactivate security tags using pointers
// func changeTagStatus(myArray *SecurityTags, tag bool) {
// 	myArray.tagStatus = tag
// }

// // Create a checkout() function which can deactivate all tags in a slice
// func checkout(myArray *[4]SecurityTags) {
// 	each := 0
// 	for each < len(myArray) {
// 		myArray[each].tagStatus = false
// 		each++
// 	}
// }

// func main() {
// 	// Create at least 4 items, all with active security tags
// 	item1 := SecurityTags{"cat", true}
// 	item2 := SecurityTags{"rat", true}
// 	item3 := SecurityTags{"cheeta", true}
// 	item4 := SecurityTags{"hen", true}
// 	// Store them in a slice or array
// 	myArray := [4]SecurityTags{item1, item2, item3, item4}
// 	fmt.Println("myArray is :  ", myArray)

// 	// Deactivate any one security tag in the array/slice
// 	changeTagStatus(&myArray[0], false)
// 	fmt.Println("changeTagStatus at myArray[0] ", myArray)

// 	// Create a checkout() function which can deactivate all tags in a slice
// 	checkout(&myArray)
// 	fmt.Println("checkout result at false all", myArray)

// }

// //--Summary:
// //  Create a program that can store a shopping list and print
// //  out information about the list.
// //
// //--Requirements:
// //* Using an array, create a shopping list with enough room
// //  for 4 products
// //  - Products must include the price and the name
// //* Insert 3 products into the array
// //* Print to the terminal:
// //  - The last item on the list
// //  - The total number of items
// //  - The total cost of the items
// //* Add a fourth product to the list and print out the
// //  information again

// package main

// type Product struct {
// 	name  string
// 	price int
// }

// var ShoppingList [5]Product

// func getTotalCost(ShoppingList [5]Product) int {
// 	sumOfPrices := 0
// 	for i := 0; i < len(ShoppingList); i++ {
// 		sumOfPrices += ShoppingList[i].price
// 	}
// 	return sumOfPrices
// }

// func main() {

// 	pc0 := Product{"apple", 11000}

// 	ShoppingList[0] = pc0

// 	pc1 := Product{"hp", 19000}

// 	ShoppingList[1] = pc1

// 	pc2 := Product{"lenovo", 13000}

// 	ShoppingList[2] = pc2

// 	fmt.Println("Last_item on list", ShoppingList[len(ShoppingList)-1])
// 	fmt.Println("total num of items on list is ", len(ShoppingList))
// 	fmt.Println("total cost of items on list", getTotalCost(ShoppingList))

// 	pc3 := Product{"dell", 11560}

// 	ShoppingList[3] = pc3

// 	pc4 := Product{"lg", 13400}

// 	ShoppingList[4] = pc4
// 	fmt.Println("Some items added to the List")
// 	fmt.Println("Last_item on list", ShoppingList[len(ShoppingList)-1])
// 	fmt.Println("total num of items on list is ", len(ShoppingList))
// 	fmt.Println("total cost of items on list", getTotalCost(ShoppingList))
// 	// getTenNums()
// }

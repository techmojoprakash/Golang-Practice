package main

import "fmt"

func main() {

	// create map with values
	newMap := map[string]int{"cat": 1, "cow": 2}

	// create map using make
	// map is based on Hash Function, so it's not in order
	// unable to access through index
	myMap := make(map[string]int)

	// assign values to myMap
	myMap["A"] = 980
	myMap["B"] = 344
	myMap["C"] = 543
	myMap["D"] = 651
	myMap["E"] = 124
	myMap["E"] = 999            // replace key E with given value
	myMap["F"] = myMap["F"] + 1 // get value and increment and put same place
	// if G not present then it will return 0 and we are adding 1 to it and place it in G key
	myMap["G"] = myMap["F"] + 1

	fmt.Println("myMap :=> ", myMap)

	// loop map byusing range function
	// inplace of key or map you can use _
	for key, value := range newMap {
		fmt.Printf("at newMap=> key is %v and value is %v\n", key, value)
	}

	// get all values sum
	sum := 0
	for _, value := range myMap {
		sum += value
	}

	fmt.Println("at myMap :=> total values sum is :=> ", sum)

	// muting maps

	// delete
	delete(myMap, "B")
	fmt.Println("B deleted then updated myMap :=> ", myMap)

	//retrive map value
	c_data := myMap["C"]
	fmt.Println("C data is :=> ", c_data)

	// test key are present or not
	// membership test
	// flag hold value of True or False
	// if value present then flag is true else flase
	dataVal, flag := myMap["C"]
	if flag {
		fmt.Println("Data is Present", dataVal)
	} else {
		fmt.Println("Data is NOT Present", flag)

	}

}

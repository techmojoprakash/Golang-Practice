package main

import "fmt"

func main() {
	fmt.Println("Hello")
	myArray := [10]int{6, 7, 8, 9, 10}

	// method - 1
	for i := 1; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	// method - 2 ( acts as While loop )
	i := 0
	for i < len(myArray) {
		fmt.Println(myArray[i])
		i++
	}

	// method - 3 Infinite loop
	j := 0
	for {
		banty := myArray[j]
		if banty == 0 {
			break
			// start motor
		}
		fmt.Println(banty)
		j++

	}

	// Loops pending
	// https://www.golangprograms.com/for-range-loops.html

	// For-each range loop
	// Looping over elements in slices, arrays, maps, channels or strings is often better done with a range loop.

	// strings := []string{"hello", "world"}
	// for i, s := range strings {
	// 	fmt.Println(i, s)
	// }
}

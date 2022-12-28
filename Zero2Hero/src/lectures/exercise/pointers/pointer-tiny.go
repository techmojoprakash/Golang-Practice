// package main

// import "fmt"

// // structure of Counter
// type Counter struct {
// 	hits int
// }

// // function of Counter in which increment hits by 1

// func increment(counter *Counter) {
// 	counter.hits += 1
// }

// // 	// replacing the value of old with new and count + 1
// func replace(old *string, new string, counter *Counter) {
// 	// accepting pointer and dereferenced and increment value by 1
// 	// simply indirectly updating the actual value based on it's address
// 	*old = new
// 	increment(counter)
// 	fmt.Println(counter.hits)
// }

// func main() {
// 	fmt.Println("Hello")
// 	first := "welcome"
// 	last := "Techmojo"
// 	repWord := "Jayaprakash"
// 	counter := Counter{}
// 	fmt.Println(first, last, repWord, counter)
// 	fmt.Println("actual first value is ", first, "actual count value is ", counter.hits)

// 	// replacing the value of first with repWord and count + 1
// 	replace(&first, repWord, &counter)
// 	fmt.Println("updated first value is ", first, "updated count value is ", counter.hits)

// 	myArray := [2]string{"TCS", "Wipro"}
// 	fmt.Println("actual myArray[0] value is ", myArray[0], "actual count value is ", counter.hits)

// 	// replacing the value of myArray[0] with last and count + 1
// 	replace(&myArray[0], last, &counter)
// 	fmt.Println("updated myArray[0] value is ", myArray[0], "updated count value is ", counter.hits)

// }

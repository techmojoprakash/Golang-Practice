package main

import (
	"fmt"
	"sync"
)

func printSomething(inputTxt string, wg *sync.WaitGroup) {
	fmt.Println(inputTxt)
	defer wg.Done()
}

// Note : Wait group should always pass as pointer but not as copy
func main() {
	var wg sync.WaitGroup
	numbersList := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
	}

	wg.Add(len(numbersList))
	for i, item := range numbersList {
		go printSomething(fmt.Sprintf("%d : %s", i, item), &wg)

	}
	wg.Wait()
	fmt.Println("*********Done*********")
}

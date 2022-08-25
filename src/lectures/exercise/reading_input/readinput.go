package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Go Program started")

	// bufio is a package to read user inputs
	// os is an pkg in golang , which can accepts with inputs from user
	myReader := bufio.NewReader(os.Stdin)

	// PART-1
	// we don't have try cache blocks in golang
	// we can use comma ok insted
	// data, err = expression
	// above one is syntax
	// data = output of expression
	// err = if expression throws an error , this variable cache the error
	// in ReadString , you shound not use double quotes, use single quotes only

	// fmt.Println("Please enter input : ")
	// input, err := myReader.ReadString('\n') // Reading string
	// // input, _ := myReader.ReadString('\n')     // err or input can be ignored by using _

	// fmt.Println("given input is : ", input, " error is : ", err)

	// PART-2
	// Reading input as integer
	// we didn't have saparate function in bufio pkg
	// so we can use strconv pkg ParseInt function
	// again we will get "\n" unknown char at input
	// so we can use "strings" pkg to trim leading and trailing spaces while reading input

	fmt.Println("Please enter input : ")
	raw_input, _ := myReader.ReadString('\n')
	input, err := strconv.ParseInt(strings.TrimSpace(raw_input), 0, 64)
	fmt.Println("input is : ", input, ", of type ", reflect.TypeOf(input), ", error is: ", err)

}

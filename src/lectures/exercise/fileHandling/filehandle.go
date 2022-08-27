package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// reading file
func readfileNow(filename string) {
	// while reading data from file it would be in bytes format

	databytes, err := ioutil.ReadFile(filename)
	checkNilError(err)

	// we need to convert the data in bytes format to string
	fmt.Println("Text data inside the file is \n", string(databytes))

}

func checkNilError(err error) {
	if err != nil {
		// this panic keyword simply print that error and exit
		panic(err)
	}
}
func main() {

	fmt.Println("File handiling demo")

	// adding content to variable
	content := "This is data , we are going to save it in file"
	// checkNilError(err)
	// creating file
	file, err := os.Create("./myfile.txt")
	checkNilError(err)
	//  writing content to created file
	length, err := io.WriteString(file, content)
	checkNilError(err)
	// printing length
	fmt.Println("lenght,", length)
	// reading the file
	readfileNow("myfile.txt")

}

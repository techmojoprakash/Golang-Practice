package main

import (
	"flag"
	"fmt"
)

var (
	// Defining a string flag
	strFlag = flag.String("language", "Golang", "Golang is the awesome google language")
	// Defining an integer flag
	intFlag = flag.Int("downloads", 1000000, "Number of times Go has been downloaded")
	// Defining a boolean flag
	boolFlag = flag.Bool("isAwesome", true, "Yes! Go is awesome")
)

func main() {

	// Call flag.Parse() to parse the command-line flags
	flag.Parse()
	// Log the flags to the terminal
	fmt.Println("String flag ",
		*strFlag,
		*intFlag)
	fmt.Println("Integer flag ", *intFlag)
	fmt.Println("Boolean flag ", *boolFlag)
}

// Run
// go run flags.go -language=Python -downloads=2000000 -isAwesome=true

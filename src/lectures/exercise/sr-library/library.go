//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

//	type Book struct {
//		nameOfBook string
//	}
type TMLibrary struct {
	book          string
	nameOfStudent string
	// BookLog map
	// timeOfCheckout time.Time
	// timeOfReturned time.Time
}

func BookCheckoutFromLibrary(BookLibrary map[string]bool, bookName string, BookLog map[time.Time]string, studentNameVal string) {
	if BookLibrary[bookName] {
		BookLibrary[bookName] = false
		logItem := [2]string{bookName, studentNameVal}
		BookLog[time.Now()] = logItem
		// StudentLibrary
		fmt.Println(bookName, " Book Checkout successfully From Library")
	} else {
		fmt.Printf("Sorry! %v is not available", bookName)
	}

}

func BookReturnToLibrary(BookLibrary map[string]bool, bookName string, BookLog map[time.Time]string) {
	BookLibrary[bookName] = true
	BookLog[time.Now()] = bookName
	fmt.Println(bookName, "Returned successfully To Library")

}

// adding book to BookLibrary
func AddBook2Library(BookLibrary map[string]bool, bookName string) {
	BookLibrary[bookName] = true
	fmt.Println(bookName, "added successfully To Library")

}

func AddStudent2Library(StudentLibrary map[string]string, icard string, studentName string) {
	StudentLibrary[icard] = studentName
	fmt.Println(studentName, "added successfully To StudentLibrary")

}

func main() {

	// maintain maps to store books uniquely
	BookLibrary := make(map[string]bool)
	StudentLibrary := make(map[string]string)
	BookLog := make(map[time.Time][2]string)
	// adding book to BookLibrary
	AddBook2Library(BookLibrary, "TCS Book")
	AddBook2Library(BookLibrary, "CISCO Book")
	AddBook2Library(BookLibrary, "TM Book")
	AddBook2Library(BookLibrary, "APPLE Book")
	AddBook2Library(BookLibrary, "DELL Book")

	// adding student to student library
	AddStudent2Library(StudentLibrary, "PR", "Prakash")
	AddStudent2Library(StudentLibrary, "VI", "Vinay")
	AddStudent2Library(StudentLibrary, "AR", "Arjun")
	AddStudent2Library(StudentLibrary, "NA", "Nakul")
	AddStudent2Library(StudentLibrary, "SA", "Sahadev")

	fmt.Println("BookLibrary is : ", BookLibrary)
	fmt.Println("StudentLibrary is : ", StudentLibrary)
	fmt.Println("BookLog is : ", BookLog)

	// take book from library and maintain log
	BookCheckoutFromLibrary(BookLibrary, "DELL Book", BookLog, StudentLibrary["VI"])
	fmt.Println("Current BookLog is : ", BookLog)
	time.Sleep(2 * time.Second)
	BookCheckoutFromLibrary(BookLibrary, "TM Book", BookLog)

	fmt.Println("Current BookLog is : ", BookLog)

	// return book to library and maintain log
	BookReturnToLibrary(BookLibrary, "DELL Book", BookLog)

	// fmt.Println("Current BookLibrary is : ", BookLibrary)
	// fmt.Println("Current StudentLibrary is : ", StudentLibrary)
	fmt.Println("Current BookLog is : ", BookLog)

	fmt.Println("final BookLibrary is : ", BookLibrary)
	fmt.Println("final StudentLibrary is : ", StudentLibrary)
}

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

type Library struct {
	student []string
	books   map[string]bool
	booklog []EachBookLog
}

type EachBookLog struct {
	bookName       string
	studentName    string
	checkOutReturn string //checkout or return
	time           string
}

func BookCheckOut(mylib *Library, eachbklog *EachBookLog, bkName string, stdName string) {
	// checking book availability in Library
	if mylib.books[bkName] {
		// making book unavailable
		mylib.books[bkName] = false
		fmt.Println("book making unavailable done : ", mylib.books[bkName])
		// creating eachbook struct with all fields
		eachbklog.bookName = bkName
		eachbklog.checkOutReturn = "checkout"
		eachbklog.studentName = stdName
		eachbklog.time = time.Now().String()
		fmt.Println("creating eachbook struct with all fields done : ", eachbklog)

		// pushing eachbook struct to booklog slice
		mylib.booklog = append(mylib.booklog, *eachbklog)
		fmt.Println("pushing eachbook struct to booklog slice done : ", mylib.booklog)
	} else {
		fmt.Println("Sorry Book is not available")
	}
}

func BookReturn(mylib *Library, eachbklog *EachBookLog, bkName string, stdName string) {
	// making book unavailable
	mylib.books[bkName] = true
	fmt.Println("book making available done : ", mylib.books[bkName])
	// creating eachbook struct with all fields
	eachbklog.bookName = bkName
	eachbklog.checkOutReturn = "return"
	eachbklog.studentName = stdName
	eachbklog.time = time.Now().String()
	fmt.Println("creating eachbook struct with all fields done : ", eachbklog)

	// pushing eachbook struct to booklog slice
	mylib.booklog = append(mylib.booklog, *eachbklog)
	fmt.Println("pushing eachbook struct to booklog slice done : ", mylib.booklog)

}

func main() {
	var TMLib Library
	var eachbkLog EachBookLog

	// adding students to library
	TMLib.student = append(TMLib.student, "Jai", "Kim", "Sai", "Siva", "Suhas")
	fmt.Println("TMLib.student ", TMLib.student)

	TMLib.books = make(map[string]bool)
	TMLib.books["TCS Book"] = true
	TMLib.books["Wipro Book"] = true
	TMLib.books["Apple Book"] = true
	TMLib.books["DELL Book"] = true
	TMLib.books["Deloit Book"] = true
	TMLib.books["Adani Book"] = true
	fmt.Println("TMLib.books ", TMLib.books)

	fmt.Println("Current Library \n", TMLib)

	// eachbklog := EachBookLog("TCS Book", "Jai", "checkout", getTime())
	// book checkout
	// params - bookLibrary, EachBookLog, bookname, studentName
	BookCheckOut(&TMLib, &eachbkLog, "TCS Book", TMLib.student[0])
	BookCheckOut(&TMLib, &eachbkLog, "TCS Book", TMLib.student[1])
	fmt.Println("Current Library \n", TMLib)
	fmt.Println("-----------------------------------------")
	BookReturn(&TMLib, &eachbkLog, "TCS Book", TMLib.student[0])

	fmt.Println("Current Library \n", TMLib)

}

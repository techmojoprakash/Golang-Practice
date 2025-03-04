/*

University will release marks to ECE and MECH students.
Here University has list of marks then it will dispatch the marks to
students based on branch - ie. it will forward markssheet to branch channels

branches receptors run in go routine to receive the data from respective branch channel



*/

package main

import (
	"fmt"
	"sync"
)

type MarksSheet struct {
	Id     int
	Branch string
	Marks  int
}

var wg sync.WaitGroup

// ece_chan <- *eachSheet even we are sending pointer to channel, but at the receiver end it will the copy of object .

// This method will dispatch the marks sheet based on Branch to specific channel
func CentralBoard(eachSheet *MarksSheet, mech_chan, ece_chan chan MarksSheet) {
	// defer wg.Done()

	switch eachSheet.Branch {
	case "ECE":
		fmt.Printf("Sending 2 ECE Channel id is %d\n", eachSheet.Id)
		ece_chan <- *eachSheet

	case "MECH":
		fmt.Printf("Sending 2 MECH Channel id is %d\n", eachSheet.Id)
		mech_chan <- *eachSheet

	}
}

func GetECEMarks(ece_chan chan MarksSheet) {
	defer wg.Done()
	fmt.Println("GetECEMarks func Started....!")
	for {
		ece_sheet, ok := <-ece_chan
		if !ok {
			fmt.Println("ECE Channel is closed...")
			return
		} else {
			fmt.Println("we received the ECE marks sheet", ece_sheet)
		}
	}
}

func GetMECHMarks(mech_chan chan MarksSheet) {
	defer wg.Done()
	fmt.Println("GetMECHMarks func Started....!")
	for {
		mech_sheet, ok := <-mech_chan
		if !ok {
			fmt.Println("MECH Channel is closed...")
			return
		} else {
			fmt.Println("we received the MECH marks sheet", mech_sheet)
		}
	}
}

func ReleaseMarks(mech_chan, ece_chan chan MarksSheet) {
	All_Marks := []MarksSheet{
		{Id: 901, Branch: "MECH", Marks: 67},
		{Id: 902, Branch: "MECH", Marks: 93},
		{Id: 903, Branch: "MECH", Marks: 47},
		{Id: 904, Branch: "MECH", Marks: 71},
		{Id: 905, Branch: "MECH", Marks: 90},
		{Id: 501, Branch: "ECE", Marks: 28},
		{Id: 502, Branch: "ECE", Marks: 81},
		{Id: 503, Branch: "ECE", Marks: 48},
		{Id: 504, Branch: "ECE", Marks: 62},
		{Id: 505, Branch: "ECE", Marks: 99},
		{Id: 506, Branch: "ECE", Marks: 10},
	}
	// wg.Add(len(All_Marks))
	for i, eachSheet := range All_Marks {
		// Based on branch marks sheet should send to specific channel
		fmt.Printf("Releasing eachSheet to Central Board, Index : %d , Sheet Data : %d and Branch %s \n", i, eachSheet.Id, eachSheet.Branch)
		// Note : Initially CentralBoard was in goroutine, that leads to deadlock issues. then removed.
		CentralBoard(&eachSheet, mech_chan, ece_chan)
	}
	close(mech_chan)
	close(ece_chan)
}

func main() {

	fmt.Println("Multi Channel Comm - One to Many")

	mech_chan := make(chan MarksSheet)
	ece_chan := make(chan MarksSheet)
	wg.Add(2)
	go GetECEMarks(ece_chan)
	go GetMECHMarks(mech_chan)

	ReleaseMarks(mech_chan, ece_chan)
	wg.Wait()
	fmt.Println("DONE with Main func.....!")
}

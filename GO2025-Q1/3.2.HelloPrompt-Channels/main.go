package main

import "fmt"

func main() {
	fmt.Println("Main")
	TICKET_ALIAS := "TICKET_"
	myid := "BT-123"
	bookMakerID := 323232
	result := TICKET_ALIAS + fmt.Sprint(myid) + "_" + fmt.Sprint(bookMakerID)
	fmt.Println("result", result)
}

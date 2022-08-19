package main

import "fmt"

type Passenger struct {
	name      string
	ticketNum int
	boarded   bool
}

type bus struct {
	FrontPassenger Passenger
}

func main() {
	fmt.Println("Welcome to bus program")
	vinay := Passenger{"vinay", 12345, true}
	myBus := bus{vinay}
	fmt.Println("Front Passenger name is : ", myBus.FrontPassenger.name)

	var (
		chandu   = Passenger{"chandu", 12346, false}
		anurag   = Passenger{"anurag", 12347, false}
		srikanth = Passenger{"srikanth", 12348, false}
	)

	fmt.Println("all passenger names are", vinay.name, chandu.name, anurag.name, srikanth.name)

	// board one passenger
	chandu.boarded = true
	myBus = bus{chandu}
	fmt.Println("Front Passenger name is : ", myBus.FrontPassenger.name)

	anurag.ticketNum = 45632
	fmt.Println("Anurag ticket number is : ", anurag.ticketNum)

	srikanth.name = "Srikanth A"
	fmt.Println("Anurag new name is : ", srikanth.name)

	// board one passenger
	srikanth.boarded = true
	myBus = bus{srikanth}
	fmt.Println("Front Passenger name is : ", myBus.FrontPassenger.name)

}

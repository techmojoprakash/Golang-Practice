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

// method-1
// related to pointers concept
// we are not passing struct in parameters
// person is variable
// *Passenger is pointer
// this model is golang specific
// you can call this from all Pasenger structures
// for example chandu is one passenger structure then
// chandu.changeName() you should call in this way
// then structure is passed and nameval will be updated as per logic

func (person *Passenger) changeName(nameval string) {
	// Pointer to a Struct in Golang
	// using a pointer but here
	// we are not using dereferencing explicitly
	// person.name = nameval

	// another way , dereferencing explicitly
	(*person).name = nameval

	fmt.Println("changeName", person.name)
}

// //method-2
// related to pointers concept
// // we are passing struct in parameters
// person is variable
// *Passenger is pointer
// this model is general like python, java

// func changeName(person *Passenger, nameval string) {
// 	person.name = nameval
// 	fmt.Println("changeName", person.name)
// }

func main() {
	fmt.Println("Welcome to bus program")
	p1 := Passenger{"vinay", 12345, true}
	myBus := bus{p1}
	fmt.Println("Front Passenger name is : ", myBus.FrontPassenger.name)

	var (
		chandu   = Passenger{"chandu", 12346, false}
		anurag   = Passenger{"anurag", 12347, false}
		srikanth = Passenger{"srikanth", 12348, false}
	)

	fmt.Println("all passenger names are", p1.name, chandu.name, anurag.name, srikanth.name)

	// board one passenger
	chandu.boarded = true
	myBus = bus{chandu}
	fmt.Println("Front Passenger name is : ", myBus.FrontPassenger.name)

	anurag.ticketNum = 45632
	fmt.Println("Anurag ticket number is : ", anurag.ticketNum)

	// srikanth.name = "Srikanth A"

	// method - 1
	srikanth.changeName("Srikanth A")
	// method-2
	// we should pass reference ex: &srikanth, &chandu
	// starts with & symbol
	// changeName(&srikanth, "Srikanth A")

	fmt.Println("srikanth new name is : ", srikanth.name)

	// board one passenger
	srikanth.boarded = true
	myBus = bus{srikanth}
	fmt.Println("Front Passenger name is : ", myBus.FrontPassenger.name)

	// Printf vs Println in structs
	// printf : %+v : will print variable name with values
	// println : will print values only

	fmt.Println("srikanth=> ", srikanth)
	fmt.Printf("srikanth=> %+v ", srikanth)

}

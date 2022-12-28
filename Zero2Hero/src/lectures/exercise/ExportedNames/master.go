package main

import (
	"fmt"
	"mypack/firstTM"
	"mypack/second"
)

func main() {
	firstData := firstTM.GetName("TechMojo")
	fmt.Println(firstData)
	secondData := second.GetVillageName("Hyderabad")
	fmt.Println(secondData)

}

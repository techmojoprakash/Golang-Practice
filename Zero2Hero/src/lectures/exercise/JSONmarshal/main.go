package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Ref : https://www.youtube.com/watch?v=ZM5tlKa-iW8
// Marshal can convert data in to bytes data , unmarshal vice versa
// Marshal => Struct to JSON
// Unmarshal => JSON to Struct

type Person struct {
	Name string `json:"name"` // if you change "name" here , json out put will change
	Age  int    `json:"age"`
}

func main() {

	//Marshal
	p := Person{Name: "Jayaprakash"} // creating struct
	// as we know that if we didn't specify value while creating struct ,
	//  then struct will assign default values to it
	pBytes, err := json.Marshal(p)
	// pBytes is bytes data
	log.Print(err) //logger is new package , we can use like printf
	log.Print(string(pBytes))
	fmt.Printf("pBytes type is : %T", pBytes) // type

	// Unmarshal
	RawJSON := json.RawMessage(`{"name": "prakash", "Age": 23}`)
	var New_Person Person
	err2 := json.Unmarshal(RawJSON, &New_Person)
	log.Print(err2)
	log.Print(New_Person)
	fmt.Printf("pBytes type is : %T", New_Person) // type
}

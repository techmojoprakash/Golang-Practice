package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("HelloWorld")
	fmt.Println("Server Started at 8080 port")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error is : ", err)
		return
	}
}

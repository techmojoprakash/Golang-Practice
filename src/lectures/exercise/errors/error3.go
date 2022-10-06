// // Custom Errors in Golang

// // In Go, we can create custom errors by implementing an error interface in a struct.

// package main

// import (
// 	"fmt"
// 	"os"
// )

// type MyError struct{}

// type error interface {
// 	Error() string
// }

// func (m *MyError) Error() string {
// 	return "boom"
// }

// func sayHello() (string, error) {
// 	return "", &MyError{}
// }

// func main() {
// 	s, err := sayHello()
// 	if err != nil {
// 		fmt.Println("unexpected error: err:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("The string:", s)
// }

// --------------------Another one--------------------

// Collecting Detailed Information in a Custom Error

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// type RequestError struct {
// 	StatusCode int

// 	Err error
// }

// func (r *RequestError) Error() string {
// 	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
// }

// func doRequest() error {
// 	return &RequestError{
// 		StatusCode: 503,
// 		Err:        errors.New("unavailable"),
// 	}
// }

// func main() {
// 	err := doRequest()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("success!")
// }

// ----output------

// Output
// status 503: err unavailable
// exit status 1

// ref : https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
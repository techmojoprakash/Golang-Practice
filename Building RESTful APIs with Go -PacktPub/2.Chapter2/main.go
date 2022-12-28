package main

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Running API with This Homepage at Specified port"))
	if err != nil {
		return
	}

}

func main() {
	fmt.Println("Chapter - 2 Started")
	http.HandleFunc("/", HomePage)
	fmt.Println("Server Started at 8080 port")
	fmt.Println("HIT ==> curl localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

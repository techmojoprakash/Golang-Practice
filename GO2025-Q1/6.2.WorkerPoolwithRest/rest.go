package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var server http.Server

type Task struct {
	Id   int
	Data int
}

type TaskWorker struct {
	Task
	W http.ResponseWriter
}

type Resp struct {
	Id   int `json:"id"`
	Data int `json:"data"`
}

// REST HandlerFunc
func sendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		fmt.Println("Only Post Method Allowed")
		http.Error(w, "Only Post Method Allowed", http.StatusMethodNotAllowed)
	}
	task := new(Task)
	// json.Unmarshal([]byte(r.Body), &task)
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println("Decode Error")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println("INPUT", task)

	tw := TaskWorker{Task: *task, W: w}

	// feed the http decoded data to TaskCh Channel, so worker will read this channel
	TaskCh <- &tw
	// Once Worker completed processing then it will push to ResultCh , here we are capturing data
	res := <-ResultCh
	fmt.Printf("ResultCh Data Received %v\n", res)
	res.W.Header().Set("Content-Type", "application/json")
	// Encode to JSON *here* in the worker
	respBody, err := json.Marshal(res.Apiresp) // Get the JSON as a byte slice
	res.W.Write(respBody)
	if err != nil {
		http.Error(res.W, err.Error(), http.StatusInternalServerError)
		fmt.Println("Encoding Error", err.Error())
	}
	fmt.Println("Response Sent for request", string(respBody))
}

func InitREST() {
	fmt.Println("Hello REST...")
	mux := http.NewServeMux()

	server = http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	mux.HandleFunc("/send", sendData)
	fmt.Println("REST server started at port :", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Error", err)
		}
	}()

}

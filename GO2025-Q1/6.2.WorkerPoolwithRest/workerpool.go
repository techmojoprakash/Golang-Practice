package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type APIResp struct {
	Id      int
	Data    int
	Comment string
}

type Result struct {
	W       http.ResponseWriter
	Apiresp APIResp
}

// Worker
// it will read the data from taskCh and process and then send resp to ResultCh
func restworker(workerID int, taskCh chan *TaskWorker, resultCh chan Result) {
	defer wg.Done()
	for eachTaskW := range taskCh {
		fmt.Printf("WorkerID %d Task is %v", workerID, eachTaskW)

		result := Result{Apiresp: APIResp{Id: eachTaskW.Id, Data: eachTaskW.Data, Comment: "Thanks"},
			W: eachTaskW.W}

		fmt.Printf("Worker %d Processed the data", workerID)
		resultCh <- result
	}

}

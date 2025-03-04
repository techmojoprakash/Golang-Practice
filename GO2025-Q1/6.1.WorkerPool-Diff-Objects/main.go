package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Task struct {
	id   int
	data int
}

type Result struct {
	Task
	comment string
}

func worker(WorkerID int, taskCh <-chan Task, resultCh chan<- Result) {
	defer wg.Done()
	for eachTask := range taskCh {
		fmt.Printf("worker no %d received work\n", WorkerID)
		eachTask.data = eachTask.data * 100
		res := Result{Task: eachTask, comment: fmt.Sprintf("Processed by worker %d", WorkerID)}
		fmt.Println("worker out", WorkerID, res)
		resultCh <- res
	}

}

func main() {

	t1 := time.Now()
	fmt.Println("Hello Workerpool with Diff Objects")
	workersize := 20
	taskChBuffSize := 50
	resultChBuffSize := 50

	taskCh := make(chan Task, taskChBuffSize)
	resultCh := make(chan Result, resultChBuffSize)

	// Initialise workers
	func() {
		for i := 1; i <= workersize; i++ {
			wg.Add(1)
			go worker(i, taskCh, resultCh)
		}
	}()

	allTasks := []Task{}
	// prepare Tasks
	for x := 1; x < 10; x++ {
		allTasks = append(allTasks, Task{id: x, data: x + 10})
	}
	fmt.Println("allTasks", allTasks)

	// Feed the data to TaskCh
	go func() {
		for _, each := range allTasks {
			taskCh <- each
		}
		close(taskCh)

	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Read the data from ResultCh
	for res := range resultCh {
		fmt.Printf("ResultCh Data Received %v\n", res)
	}
	fmt.Println("DONE", time.Since(t1))
}

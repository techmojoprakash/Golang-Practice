// WorkerPool Implementation : https://www.youtube.com/watch?v=7IDKae4SvDw
// https://github.com/arshad404/CodePiper/blob/main/GO/concurrency-patterns/Worker-Pool/main.go

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Task struct {
	id        int
	workernum int
}

// each worker do some processing
// read data from taskCh and process then push to resCh
func worker(workerID int, taskCh <-chan Task, resCh chan<- Task) {
	defer wg.Done()
	for j := range taskCh {
		fmt.Printf("%d Worker Received data : %v\n", workerID, j)
		out := j.id * 100
		j.id = out
		j.workernum = workerID
		fmt.Printf("%d DONE Processed data : %v\n", j.workernum, j)
		resCh <- j
	}

}
func main() {
	fmt.Println("WorkerPools....!")
	workerSize := 5
	taskChQueSize := 10
	resultChQueSize := 10

	tasksCh := make(chan Task, taskChQueSize)
	resCh := make(chan Task, resultChQueSize)

	//Start worker pool - create workers

	func() {
		for workerID := 1; workerID <= workerSize; workerID++ {
			wg.Add(1)
			go worker(workerID, tasksCh, resCh)
			fmt.Println("Worker created id is ", workerID)
		}
	}()

	taskList := []Task{{id: 1},
		{id: 2},
		{id: 3},
		{id: 4},
		{id: 5},
		{id: 6},
		{id: 7},
		{id: 8},
		{id: 9},
		{id: 10},
		{id: 11},
	}

	// send tasks to taskCh channel
	for _, eachTsk := range taskList {
		tasksCh <- eachTsk
	}
	close(tasksCh)

	// once workers finished all tasks then close resCh
	go func() {
		wg.Wait()
		close(resCh)
	}()

	for result := range resCh {
		fmt.Printf("Result got it, worker is %d and data is %v\n", result.workernum, result)
	}

	fmt.Println("Awesome....!")
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var TaskCh chan *TaskWorker
var ResultCh chan Result

func main() {
	// Init REST Server
	InitREST()

	t1 := time.Now()
	fmt.Println("Hello Workerpool with Diff Objects")
	workersize := 2
	taskChBuffSize := 5
	resultChBuffSize := 5

	// Channel Buffer sizes
	TaskCh = make(chan *TaskWorker, taskChBuffSize)
	ResultCh = make(chan Result, resultChBuffSize)

	// Initialise workers
	func() {
		for i := 1; i <= workersize; i++ {
			wg.Add(1)
			go restworker(i, TaskCh, ResultCh)
		}
	}()

	// In general when we generate data from fixed array, map then we will close TaskCh after for loop. then we close ResultCh after all workers done
	// but here REST , so we don't know when can we get the request so we close TaskCh after all workers done.
	go func() {
		wg.Wait()
		close(TaskCh) // Close TaskCh after all tasks are submitted
	}()

	// Keep main running until a signal is received
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // Listen for Ctrl+C or kill signal
	<-quit                                               // Block until signal received
	fmt.Println("Shutting down server...")
	fmt.Println("DONE", time.Since(t1))
}

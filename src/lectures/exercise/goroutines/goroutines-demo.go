package main

import (
	"fmt"
	"sync"
)

/*

Sync.WaitGroup  // under sync package
Wait group is helping us to run gorotines inside main gorotine
Wait can be used to block until all goroutines have finished.
waitgroup has 3 methods
1.	Add(1) – should put before gorotine &&& Add adds delta, which may be negative, to the WaitGroup counter. If the counter becomes zero, all goroutines blocked on Wait are released. If the counter goes negative, Add panics.
2.	Done() – should place inside child gorotine, it says I’m done - Done decrements the WaitGroup counter by one.
a.	Protip : use defer statement , which helps to execute done() at the end of child gorotine
3.	Wait() - Wait blocks until the WaitGroup counter is zero. In another words, it will not allow main method to finish execution. ~ it should place at end of main method


*/

// account balance deposit and debit methods implementation
// by using goroutine, mutex
var balance int = 1000
var mut sync.Mutex

func creditMoney(money int, wg *sync.WaitGroup) {
	mut.Lock()
	balance += money
	fmt.Println("after func creditMoney Current balance :", balance)
	mut.Unlock()
	wg.Done()
}

func debitMoney(money int, wg *sync.WaitGroup) {
	mut.Lock()
	balance -= money
	fmt.Println("after debitMoney Current balance :", balance)
	mut.Unlock()
	wg.Done()
}

func main() {

	fmt.Println("Goroutine example")
	fmt.Println("Current balance :", balance)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go creditMoney(500, wg)
	go debitMoney(200, wg)
	// above two gorotines will run in sync but not in proper order
	// ultimately while credit the money debit is not possible , here mutex is helping by locks
	// similarly, while debit the money credit is also not possible because mutex helping here
	// in this way we can avoid Race Condition problem

	wg.Wait()
	fmt.Println("Current balance :", balance)

}

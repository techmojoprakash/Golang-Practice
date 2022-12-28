package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func credit(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	defer wg.Done()
	balance += amount
	mutex.Unlock()
	fmt.Println("Bank Account balance after credit", balance)
}

func debit(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	defer wg.Done()
	balance -= amount
	mutex.Unlock()
	fmt.Println("Bank Account balance after debit", balance)
}

func main() {
	fmt.Println("Hello")
	balance = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go credit(100, &wg)
	go debit(50, &wg)
	wg.Wait()
	fmt.Println("Current balance:- ", balance)
}

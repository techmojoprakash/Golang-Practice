/*

One package level variable Bank_Money
where we can withdraw or deposit money in it concurently.

use Mutex lock

*/

package main

import (
	"fmt"
	"sync"
)

var Bank_Money int

var wg sync.WaitGroup
var mutex sync.Mutex

func WithDraw(m int) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("--->WithDraw: %d, current Balance :%d\n", m, Bank_Money)
	if m < Bank_Money {
		Bank_Money = Bank_Money - m
		fmt.Printf("After WithDraw %d , New Balance : %d\n", m, Bank_Money)
	} else {
		fmt.Printf("Failed %d from Bank Money %d \n", m, Bank_Money)
	}
}

func DepositMoney(m int) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("->DepositMoney: %d, current Balance :%d\n", m, Bank_Money)
	Bank_Money = Bank_Money + m
	fmt.Printf("OK DepositMoney %d to New Balance %d \n", m, Bank_Money)
}

func main() {
	fmt.Println("Hello Mutex Example")
	Bank_Money = 1000
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 1000; i = i + 5 {
			DepositMoney(i)
		}
	}()

	go func() {
		defer wg.Done()
		for j := 1; j < 1000; j = j + 10 {
			WithDraw(j)
		}
	}()
	wg.Wait()
	fmt.Println("DONE")
}

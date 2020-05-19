package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

// func Deposit(amount int) {
// 	mu.Lock()
// 	balance += amount
// 	mu.Unlock()
// }

// func Balance() int {
// 	mu.Lock()
// 	b := balance
// 	mu.Unlock()
// 	return b
// }

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	Deposit(100)
	fmt.Println(Balance())
	Deposit(100)
	fmt.Println(Balance())
	Deposit(100)
	fmt.Println(Balance())
	Deposit(100)
	fmt.Println(Balance())
}

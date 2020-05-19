package main

import "fmt"

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}
func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
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

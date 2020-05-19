package main

import "fmt"

var deposites = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposites <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposites:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
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

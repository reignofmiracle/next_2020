package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	count := 0
	for {
		fmt.Println(<-squares)
		count++
		if count > 100 {
			break
		}
	}
}

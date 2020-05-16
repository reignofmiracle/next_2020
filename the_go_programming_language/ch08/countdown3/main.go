package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fire := make(chan struct{})

	go func() {
		tick := time.Tick(1 * time.Second)
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
			<-tick
		}
		fire <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")

	select {
	//case <-time.After(10 * time.Second):
	case <-fire:
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func launch() {
	fmt.Println("launch")
}

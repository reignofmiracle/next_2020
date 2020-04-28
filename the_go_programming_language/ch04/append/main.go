package main

import (
	"fmt"

	"./my_append"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	ret := my_append.AppendInt(s, 13)
	fmt.Printf("%v\n", ret)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = my_append.AppendInt(x, i)
		fmt.Printf("%d cap%d\t%v\n", i, cap(y), y)
		x = y
	}
}

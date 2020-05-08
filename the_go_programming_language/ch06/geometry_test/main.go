package main

import (
	"fmt"

	"../geometry"
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}

func test2() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}

func test3() {
	p := geometry.Point{1, 2}
	fmt.Println(p)
	p.Test1()
	fmt.Println(p)
	p.Test2()
	fmt.Println(p)
}

package main

import (
	"fmt"
	"image/color"

	"../coloredpoint"
	"../geometry"
)

func main() {
	test1()
	test2()
}

func test1() {
	var cp coloredpoint.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
}

func test2() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredpoint.ColoredPoint{geometry.Point{1, 1}, red}
	var q = coloredpoint.ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

}

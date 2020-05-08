package geometry

import "math"

// Point ...
type Point struct{ X, Y float64 }

// Distance ...
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance ...
func (p Point) Distance(q Point) float64 {
	return Distance(p, q)
}

// ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Test1 ...
func (p *Point) Test1() {
	p.X += 10
}

// Test2 ...
func (p Point) Test2() {
	p.X += 10
}

// Path ...
type Path []Point

// Distance ...
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func newPoint(x, y float64) *point {
	return &point{x, y}
}

func (p *point) dist(dp *point) float64 {
	return math.Sqrt(math.Pow(p.x-dp.x, 2) + math.Pow(p.y-dp.y, 2))
}

func main() {
	p1 := newPoint(1.1, 2.2)
	p2 := newPoint(2.1, 3.2)

	fmt.Println(p1.dist(p2))
}

package internal

import "fmt"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p *Point) ToString() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

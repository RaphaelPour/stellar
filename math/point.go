package math

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	p.X += other.X
	p.Y += other.Y
	return p
}

func (p Point) Sub(other Point) Point {
	p.X -= other.X
	p.Y -= other.Y
	return p
}

func (p Point) String() string {
	return fmt.Sprintf("%d/%d", p.X, p.Y)
}

func (p Point) ManhattanDist(other Point) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

func (p Point) EuclideanDist(other Point) float64 {
	return math.Sqrt(float64(Pow(Abs(p.X-other.X), 2) + Pow(Abs(p.Y-other.Y), 2)))
}

func (p Point) Equal(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) Min(other Point) Point {
	if p.X > other.X {
		p.X = other.X
	}

	if p.Y > other.Y {
		p.Y = other.Y
	}
	return p
}

func (p Point) Max(other Point) Point {
	if p.X < other.X {
		p.X = other.X
	}

	if p.Y < other.Y {
		p.Y = other.Y
	}
	return p
}

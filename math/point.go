package math

type Point[T Number] struct {
	X, Y T
}

func (p Point[T]) Add(other Point[T]) Point[T] {
	p.X += other.X
	p.Y += other.Y
	return p
}

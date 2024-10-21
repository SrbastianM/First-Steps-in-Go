package square

import "fmt"

// Interface: when a lot of structures or functions have the same values. Is more efficiency use interfaces
type figure2D interface {
	Area() float64
}

type square struct {
	base float64
}
type rectangle struct {
	base   float64
	height float64
}

func (c square) Area() float64 {
	return c.base * c.base
}

func (r rectangle) Area() float64 {
	return r.base * r.height
}

func calculate(f figure2D) {
	fmt.Println("Area: ", f.Area())
}

func Square() {
	newSquare := square{base: 2}
	newRectangle := rectangle{base: 2, height: 4}

	calculate(newSquare)
	calculate(newRectangle)
}

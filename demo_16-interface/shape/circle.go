package shape

import "math"

type circle struct {
	Redius float64
}

func NewCircle(redius float64) circle {
	return circle{
		Redius: redius,
	}
}

func (c circle) area() float64 {
	// Pi * r(กำลัง 2)
	return math.Pi * math.Pow(c.Redius, 2)
}

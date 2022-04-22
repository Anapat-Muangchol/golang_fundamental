package shape

type rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width float64, height float64) rectangle {
	return rectangle{
		Width:  width,
		Height: height,
	}
}

func (r rectangle) area() float64 {
	return r.Width * r.Height
}

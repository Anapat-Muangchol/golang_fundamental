package shape

import (
	"fmt"
	"reflect"
)

type shape interface {
	area() float64
}

func GetArea(s shape) float64 {
	return s.area()
}

func ShowInfo(s shape) {
	typeName := reflect.TypeOf(s).Name()
	switch typeName {
	case "rectangle":
		r := s.(rectangle) // cast type
		fmt.Printf("Rectangle width: %.2f, height: %.2f", r.Width, r.Height)
	case "circle":
		c := s.(circle) // cast type
		fmt.Printf("Circle radius: %.2f", c.Redius)
	}
}

func CastToRectangle(s shape) {
	r, ok := s.(rectangle)
	if !ok {
		fmt.Println("Error cast to rectangle")
	} else {
		fmt.Printf("Casting to Rectangle Success | width: %.2f, height: %.2f\n", r.Width, r.Height)
	}
}

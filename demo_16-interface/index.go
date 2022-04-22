package main

// command : go mod init demo16
import (
	"demo16/shape"
	"fmt"
)

func main() {
	r := shape.NewRectangle(10, 20)
	shape.ShowInfo(r)
	fmt.Printf(" | area: %.2f\n", shape.GetArea(r))

	c := shape.NewCircle(10)
	shape.ShowInfo(c)
	fmt.Printf(" | area: %.2f\n", shape.GetArea(c))
}

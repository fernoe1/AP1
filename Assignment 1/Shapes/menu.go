package Shapes

import (
	"fmt"

	"github.com/fernoe1/Assignment1_TemirlanSapargali/Shapes/model"
)

func Start() {
	shapes := make([]model.Shape, 0, 4)
	shapes = append(shapes, model.Rectangle{
		Width:  5,
		Height: 5,
	})

	shapes = append(shapes, model.Circle{
		Radius: 5,
	})

	shapes = append(shapes, model.Square{
		Side: 5,
	})

	shapes = append(shapes, model.Triangle{
		A: 5,
		B: 5,
		C: 5,
	})

	for _, shape := range shapes {
		fmt.Printf("Shape type: %T, shape area: %f, shape perimeter: %f \n", shape, shape.Area(), shape.Perimeter())
	}
}

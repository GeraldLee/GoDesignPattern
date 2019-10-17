package composite

import "fmt"

type Shape interface {
	Draw(fillColor string)
}

type Triangle struct{}

func (t *Triangle) Draw(fillColor string) {
	fmt.Println("Drawing Triangle with color", fillColor)
}

type Circle struct{}

func (t *Circle) Draw(fillColor string) {
	fmt.Println("Drawing Circle with color", fillColor)
}

type Drawing struct {
	shapes []Shape
}

func (d *Drawing) Add(shape Shape){
	if d.shapes == nil {
		d.shapes = make([]Shape,0)
	}
	d.shapes = append(d.shapes, shape)
}

func (d *Drawing) Draw(fillColor string) {
	for _, shape := range d.shapes {
		shape.Draw(fillColor)
	}
}


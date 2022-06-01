package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("box limit is reached. Cannot add new shape")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	capacity := b.shapesCapacity

	if i > (capacity-1) || i < 0 {
		return nil, errors.New("index out of box capacity")
	}

	if b.shapes[i] == nil {
		return nil, errors.New("there is no shape with specified index")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)

	if err != nil {
		return nil, err
	}
	b.shapes[i] = b.shapes[len(b.shapes)-1]
	b.shapes = b.shapes[:len(b.shapes)-1]
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	replacedShape, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes[i] = shape
	return replacedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := range b.shapes {
		sum += b.shapes[i].CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for i := range b.shapes {
		sum += b.shapes[i].CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var modifiedShapes []Shape
	var circleCounter = 0
	for i := range b.shapes {
		shape := b.shapes[i]
		_, ok := shape.(*Circle)
		if !ok {
			modifiedShapes = append(modifiedShapes, shape)
		} else {
			circleCounter += 1
		}
	}
	if circleCounter == 0 {
		return errors.New("no circles in the box")
	}

	b.shapes = modifiedShapes
	return nil
}

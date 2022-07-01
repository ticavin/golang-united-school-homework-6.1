package golang_united_school_homework

import (
	"errors"
)

var (
	errMax    = errors.New("full")
	errIndex  = errors.New("index out of range")
	errCircle = errors.New("no circles")
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
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return errMax
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= b.shapesCapacity {
		return nil, errIndex
	}

	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) || i < 0 {
		return nil, errIndex
	} else {
		result := b.shapes[i]
		b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)
		return result, nil
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errIndex
	}

	res := b.shapes[i]
	b.shapes[i] = shape
	return res, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var s float64
	for _, val := range b.shapes {
		s += val.CalcPerimeter()
	}
	return s

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var s float64
	for _, val := range b.shapes {
		s += val.CalcArea()
	}
	return s

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var exist bool
	var box []Shape
	for _, v := range b.shapes {
		if shape, ok := v.(*Circle); ok {
			exist = true
			continue
		} else {
			box = append(box, shape)
		}
	}
	b.shapes = box
	if !exist {
		return errCircle
	}
	return nil

}

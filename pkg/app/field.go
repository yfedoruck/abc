package app

import (
	"image"
)

//width = 153
//height = 850 //896
type Field struct {
	area   image.Rectangle
	width  int
	height int
	NumX   int
	NumY   int
	matrix [][]image.Point
}

func NewField() Field {
	const (
		FieldMinY = 105 //93
		FieldMaxY = 1105
		FieldMinX = 20
		FieldMaxX = 520
	)
	height := FieldMaxY - FieldMinY
	width := FieldMaxX - FieldMinX
	f := Field{
		area:   image.Rect(FieldMinX, FieldMinY, FieldMaxX, FieldMaxY),
		height: height,
		width:  width,
		NumX:   width / CubeWidth,
		NumY:   height / CubeWidth,
	}
	f.Matrix()
	return f
}

func (r *Field) Matrix() {
	matrix := make([][]image.Point, r.NumX)
	for i := range matrix {
		matrix[i] = make([]image.Point, r.NumY)
	}
	//matrix := make([][]int, r.NumY)
	//for i := range matrix {
	//	matrix[i] = make([]int, r.NumX)
	//}
	r.matrix = matrix
	for i := 0; i < r.NumX; i++ {
		for j := 0; j < r.NumY; j++ {
			r.matrix[i][j] = image.Pt(i*CubeWidth, j*CubeWidth)
		}
	}
}

var (
	Figures = [7][4]int{
		{1, 3, 5, 7}, // I
		{2, 4, 5, 7}, // S
		{3, 5, 4, 6}, // Z
		{3, 5, 4, 7}, // T
		{2, 3, 5, 7}, // L
		{3, 5, 7, 6}, // J
		{2, 3, 4, 5}, // O
	}
)

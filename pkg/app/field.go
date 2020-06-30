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
	matrix [][]bool
	cntDel int
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
	matrix := make([][]bool, r.NumX)
	for i := range matrix {
		matrix[i] = make([]bool, r.NumY)
	}
	r.matrix = matrix
}

func (r *Field) Fill(fig TFig) {
	for _, point := range fig.a {
		r.matrix[point.x][point.y] = true
		if r.IsRowFull(point.y) {
			r.DeleteRow(point.y)
			r.cntDel++
		}
	}
}

func (r *Field) IsRowFull(num int) bool {
	for i := 0; i < r.NumX; i++ {
		if r.matrix[i][num] == false {
			return false
		}
	}
	return true
}

func (r *Field) DeleteRow(num int) {
	for j := num; j > 0; j-- {
		for i := 0; i < r.NumX; i++ {
			r.matrix[i][j] = r.matrix[i][j-1]
		}
	}
}

func (r *Field) FilledToTop() bool {
	for i := 0; i < r.NumX; i++ {
		if r.matrix[i][0] == true {
			return true
		}
	}
	return false
}

func (r *Field) Clear() {
	for i := 0; i < r.NumX; i++ {
		for j := 0; j < r.NumY; j++ {
			r.matrix[i][j] = false
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

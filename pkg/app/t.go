package app

import (
	"fmt"
)

type TFig struct {
	a       [4]Point
	field   Field
	stopped bool
	Type    Tetromino
}

func NewFig(field Field, num Tetromino) TFig {
	a := getFig(num)
	return TFig{
		field: field,
		Type:  num,
		a:     a,
	}
}

func (r *TFig) Rotate() {
	if r.IsNotRotated() {
		return
	}

	var p = Point{}
	p = r.a[1]
	b := r.a
	for i := 0; i < 4; i++ {
		x := r.a[i].y - p.y
		y := r.a[i].x - p.x
		r.a[i].x = p.x - x
		r.a[i].y = p.y + y
		if r.IsLimitExceed(i) {
			r.a = b
			break
		}
	}
}

func (r TFig) IsNotRotated() bool {
	return fmt.Sprintf("%s", r.Type) == "O"
}

func (r *TFig) MoveLeft() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].x--
		if r.IsMinX(i) || r.IsFilled(i) {
			r.a = b
			break
		}
	}
}

func (r *TFig) MoveRight() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].x++
		if r.IsMaxX(i) || r.IsFilled(i) {
			r.a = b
			break
		}
	}
}

func (r *TFig) FallDown(field *Field) {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].y++
		if r.IsMaxY(i) || r.IsFilled(i) {
			r.Stop()
			r.a = b
			field.Fill(*r)
			break
		}
	}
}

func (r TFig) IsFilled(i int) bool {
	return r.field.matrix[r.a[i].x][r.a[i].y] == true
}

func (r *TFig) Stop() {
	r.stopped = true
}

func (r TFig) IsStopped() bool {
	return r.stopped
}

func (r TFig) NotStopped() bool {
	return r.stopped == false
}

func (r TFig) IsMinX(i int) bool {
	return r.a[i].x < 0
}

func (r TFig) IsMaxX(i int) bool {
	return r.a[i].x >= r.field.NumX
}

func (r TFig) IsMaxY(i int) bool {
	return r.a[i].y >= r.field.NumY
}

func (r TFig) IsLimitExceed(i int) bool {
	return r.IsMinX(i) || r.IsMaxX(i) || r.IsMaxY(i) || r.IsFilled(i)
}

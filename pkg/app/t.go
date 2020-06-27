package app

import "fmt"

type TFig struct {
	a          [4]Point
	XMax, YMax int
	stopped    bool
	Type       Tetromino
}

func NewFig(XMax, YMax int, num Tetromino) TFig {
	a := getFig(num)
	return TFig{
		XMax: XMax,
		YMax: YMax,
		Type: num,
		a: a,
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
		if r.IsMinX(i) {
			r.a = b
			break
		}
	}
}

func (r *TFig) MoveRight() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].x++
		if r.IsMaxX(i) {
			r.a = b
			break
		}
	}
}

func (r *TFig) FallDown() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].y++
		if r.IsMaxY(i) {
			r.Stop()
			r.a = b
			break
		}
	}
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
	return r.a[i].x >= r.XMax
}

func (r TFig) IsMaxY(i int) bool {
	return r.a[i].y >= r.YMax
}

func (r TFig) IsLimitExceed(i int) bool {
	return r.IsMinX(i) || r.IsMaxX(i) || r.IsMaxY(i)
}

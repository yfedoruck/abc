package app

import "fmt"

type TFig struct {
	a [4]Point
}

func (r *TFig) get() {
	r.a = getFig(3)
}

func (r *TFig) Rotate() {
	var p = Point{}
	p = r.a[1]
	for i := 0; i < 4; i++ {
		x := r.a[i].y - p.y
		y := r.a[i].x - p.x
		r.a[i].x = p.x - x
		r.a[i].y = p.y + y
	}
}

func (r TFig) IsLimitExceed(N, M int) bool {
	for i := 0; i < 4; i++ {
		fmt.Println(r.a[i].x, r.a[i].y)
		if r.a[i].x < 0 || r.a[i].x >= N || r.a[i].y >= M {
			return true
		}
	}

	return false
}

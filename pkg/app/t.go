package app

type TFig struct {
	a          [4]Point
	XMax, YMax int
}

func NewFig(XMax, YMax int) TFig {
	return TFig{
		XMax: XMax,
		YMax: YMax,
	}
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

func (r *TFig) MoveLeft() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].x--
		if r.a[i].x < 0 {
			r.a = b
			break
		}
	}
}

func (r *TFig) MoveRight() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].x++
		if r.a[i].x >= r.XMax {
			r.a = b
			break
		}
	}
}

func (r *TFig) FallDown() {
	b := r.a
	for i := 0; i < 4; i++ {
		r.a[i].y++
		if r.a[i].y >= r.YMax {
			r.a = b
			break
		}
	}
}

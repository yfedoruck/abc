package app

type Window struct {
	width  int
	height int
	Scale  float64
}

func NewWindow() *Window {
	const (
		ScreenWidth  = 720  //590
		ScreenHeight = 1125 //960
		Scale        = 1  //960
	)
	return &Window{
		width:  ScreenWidth,
		height: ScreenHeight,
		Scale:  Scale,
	}
}
func (r Window) Height() int {
	return int(float64(r.height) * r.Scale)
}
func (r Window) Width() int {
	return int(float64(r.width) * r.Scale)
}

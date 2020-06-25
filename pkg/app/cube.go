package app

type Cube struct {
	side int
}

func NewCube() *Cube {
	const (
		Side = 50
	)
	cube := new(Cube)
	cube.side = Side
	return cube
}
func (r Cube) Side() int {
	return r.side
}

package app

func getFig(n Tetromino) [4]Point {
	var a [4]Point
	for i := 0; i < 4; i++ {
		a[i].x = Figures[n][i] % 2
		a[i].y = Figures[n][i] / 2
	}
	return a
}

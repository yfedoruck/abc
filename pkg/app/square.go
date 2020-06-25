package app

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

// 225x225
type Square struct {
	image  *ebiten.Image
	Width  int
	Height int
	X1     int
	X2     int
	Y1     int
	Y2     int
	sprite *ebiten.Image
}

func NewSquare() *Square {
	const (
		FieldMinY = 93
		FieldMaxY = 943
		FieldMinX = 16
		FieldMaxX = 169
	)
	img := LoadSprite("R.png")
	rec := image.Rect(0, 0, 50, 50)
	sub := img.SubImage(rec)

	//eImg, err := ebiten.NewImageFromImage(sub, ebiten.FilterNearest)
	//fail.Check(err)

	return &Square{
		image:  img,
		sprite: sub.(*ebiten.Image),
		Height: 122,
		Width:  122,
	}
}

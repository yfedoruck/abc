package app

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/yfedoruck/tetris/pkg/env"
	"github.com/yfedoruck/tetris/pkg/fail"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
)

type Image struct {
	Picture       image.Image
	EbitenPicture *ebiten.Image
}

func NewImage(file string) Image {
	var s = Image{}
	s.load(file)
	return s
}

func (r *Image) load(path string) {
	var err error
	r.EbitenPicture, r.Picture, err = ebitenutil.NewImageFromFile(FilePath(path), ebiten.FilterDefault)
	fail.Check(err)
}

func LoadSprite(path string) *ebiten.Image {
	var (
		err error
		img *ebiten.Image
	)
	img, _, err = ebitenutil.NewImageFromFile(FilePath(path), ebiten.FilterDefault)
	fail.Check(err)

	return img
}

//func TransparentPixel() *pixel.Sprite {
//	var picture = pixel.PictureDataFromImage(LoadSprite("1x1.png"))
//	return pixel.NewSprite(picture, picture.Bounds())
//}
//
//func SimpleSprite(file string) *pixel.Sprite {
//	var picture = pixel.PictureDataFromImage(LoadSprite(file))
//	return pixel.NewSprite(picture, picture.Bounds())
//}

func FilePath(path string) string {
	return env.BasePath() + filepath.FromSlash("/static/"+path)
}

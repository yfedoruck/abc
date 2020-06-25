package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yfedoruck/abc/pkg/app"
	_ "image/jpeg"
	"log"
)

func main() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	//img, _, err := image.Decode(bytes.NewReader(images.FiveYears_jpg))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//app.GophersImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	win := app.NewWindow()
	ebiten.SetWindowSize(win.Width(), win.Height())
	ebiten.SetWindowTitle("Tetris")
	ebiten.SetVsyncEnabled(false)
	if err := ebiten.RunGame(app.NewGame()); err != nil {
		log.Fatal(err)
	}
}

package app

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/yfedoruck/abc/pkg/fail"
	"time"
)

const (
	ScreenWidth  = 720.0  //590
	ScreenHeight = 1125.0 //960
	Scale        = 1      //960
	CubeWidth    = 50
	Delta        = 0.3
)

type Game struct {
	fps        <-chan time.Time
	last       float64
	count      int
	frameNum   int
	Background *ebiten.Image
	Square     *Square
	screen     *ebiten.Image
	scale      float64
	field      Field
	tick       bool
	tx         int
	ty         int
	figure     TFig
}

func NewGame() *Game {
	f := NewField()
	tf := TFig{}
	tf.get()
	g := &Game{
		frameNum:   8,
		Background: LoadSprite("background.png"),
		Square:     NewSquare(), //LoadSprite("R.png"),
		scale:      Scale,       //0.379,
		field:      f,
		last:       float64(time.Now().Second()),
		tx:         f.area.Bounds().Min.X,
		ty:         f.area.Bounds().Min.Y,
		figure:     tf,
	}
	g.Fps()
	return g
}

func (r *Game) Fps() {
	r.fps = time.Tick(time.Second / 60)
}

func (r *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (r *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (r *Game) Draw(screen *ebiten.Image) {
	r.DrawBg(screen)
	r.tickTack()
	r.DrawSquare(screen)

	<-r.fps
}

func (r *Game) tickTack() {
	now := float64(time.Now().Second())
	if now-r.last > Delta {
		if !r.tick {
			r.tick = true
		}
		r.last = now
	} else {
		r.tick = false
	}
}

func (r *Game) DrawBg(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(r.scale, r.scale)
	op.GeoM.Translate(0, 0)
	err := screen.DrawImage(r.Background, op)
	fail.Check(err)
}

func (r *Game) DrawSquare(screen *ebiten.Image) {
	r.listenXMoving()
	r.FallDown()

	r.listenRotate()
	for _, point := range r.figure.a {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(point.x*CubeWidth+r.tx), float64(point.y*CubeWidth+r.ty))
		err := screen.DrawImage(r.Square.sprite, op)
		fail.Check(err)
	}
}

func (r *Game) Rotate(screen *ebiten.Image) {
	var p = Point{}
	tf := TFig{}
	tf.get()
	p = tf.a[1]
	for _, point := range tf.a {
		x := point.y - p.y
		y := point.x - p.x
		point.x = p.x - x
		point.y = p.y + y
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(point.x*CubeWidth+r.tx), float64(point.y*CubeWidth+r.ty))
		err := screen.DrawImage(r.Square.sprite, op)
		fail.Check(err)
	}
}

func (r *Game) listenXMoving() {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		r.MoveRight()
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		r.MoveLeft()
	}
}

func (r *Game) listenRotate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		r.figure.Rotate()
	}
}

func (r *Game) MoveRight() {
	r.tx += CubeWidth
}
func (r *Game) MoveLeft() {
	r.tx -= CubeWidth
}
func (r *Game) FallDown() {
	if !r.tick {
		return
	}
	if r.figure.IsLimitExceed(r.field.NumX, r.field.NumY){
		panic("test")
		return
	}
	r.ty += CubeWidth
}

//func (r *Game) Move(screen *ebiten.Image) {
//	if r.tick == 1 {
//		r.count++
//	}
//
//	var p image.Point
//	if r.count < r.field.NumY {
//		p = r.field.matrix[0][r.count]
//		r.ty = p.Y + r.field.area.Min.Y
//	}
//}

type Point struct {
	x, y int
}

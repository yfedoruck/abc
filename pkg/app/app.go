package app

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/yfedoruck/abc/pkg/fail"
	"math/rand"
	"time"
)

const (
	ScreenWidth  = 720.0  //590
	ScreenHeight = 1125.0 //960
	Scale        = 1      //960
	CubeWidth    = 50
	Delta        = 600
	SmallDelta   = 30
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
	wall       []TFig
	delta      float64
}

func NewGame() *Game {
	f := NewField()
	fNum := rand.Intn(7) //TODO
	tf := NewFig(f.NumX, f.NumY, Tetromino(fNum))
	g := &Game{
		frameNum:   8,
		Background: LoadSprite("background.png"),
		Square:     NewSquare(), //LoadSprite("R.png"),
		scale:      Scale,       //0.379,
		field:      f,
		last:       tick(),
		tx:         f.area.Bounds().Min.X,
		ty:         f.area.Bounds().Min.Y,
		figure:     tf,
		wall:       make([]TFig, 0),
		delta:      Delta,
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
	now := tick()
	if now-r.last > r.delta {
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
	if r.figure.NotStopped() {
		r.listenXMoving()
		r.FallDown()
		r.listenRotate()
		r.listenFall()
	} else {
		r.ResetDelta()
		r.wall = append(r.wall, r.figure)
		r.SetNewFigure()
	}

	r.DrawFigure(r.figure, screen)

	for _, figure := range r.wall {
		r.DrawFigure(figure, screen)
	}
}

func (r Game) DrawFigure(figure TFig, screen *ebiten.Image) {
	for _, point := range figure.a {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(point.x*CubeWidth+r.tx), float64(point.y*CubeWidth+r.ty))
		err := screen.DrawImage(r.Square.sprite, op)
		fail.Check(err)
	}
}

func (r *Game) SetNewFigure() {
	fNum := rand.Intn(7)
	r.figure = NewFig(r.field.NumX, r.field.NumY, Tetromino(fNum))
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

func (r *Game) listenFall() {
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		r.delta = SmallDelta
	} else {
		r.ResetDelta()
	}
}

func (r *Game) ResetDelta() {
	r.delta = Delta
}

func (r *Game) MoveRight() {
	r.figure.MoveRight()
}
func (r *Game) MoveLeft() {
	r.figure.MoveLeft()
}
func (r *Game) FallDown() {
	if !r.tick {
		return
	}
	r.figure.FallDown(r.field)
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

func tick() float64 {
	return float64(time.Now().UnixNano() / 1e6)
}

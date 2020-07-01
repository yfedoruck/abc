package app

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/yfedoruck/tetris/pkg/fail"
	"golang.org/x/image/font"
	"image/color"
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
	SideDelta    = 100
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
	nextFig    TFig
	delta      float64
	font       font.Face
	isEnd      bool
}

func NewGame() *Game {
	f := NewField()
	tf := NewFig(f, RandomNum())
	err := tf.StartPosition()
	fail.Check(err)

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
		delta:      Delta,
		nextFig:    NewFig(f, RandomNum()),
		font:       FontFace(),
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
	r.screen = screen
	r.DrawBg()
	r.tickTack()

	if r.isEnd {
		r.GameOver()
		return
	} else if r.field.FilledToTop() {
		r.EndGame()
		return
	}

	r.DrawNextFigure()
	r.listenMoving()
	r.FallDown()

	r.DrawScores()
	r.DrawFigure()
	r.DrawWall()

	<-r.fps
}

func (r *Game) tickTack() {
	now := tick()
	if now-r.last > r.delta {
		if !r.tick {
			r.tick = true
		}
		r.last = now
	} else if r.tick {
		r.tick = false
	}
}

func (r *Game) tact(fn func(), delta float64) {
	now := tick()
	if now-r.last > delta {
		fn()
		r.last = now
	}
}

func (r *Game) DrawBg() {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(r.scale, r.scale)
	op.GeoM.Translate(0, 0)
	err := r.screen.DrawImage(r.Background, op)
	fail.Check(err)
}

func (r *Game) DrawWall() {
	for i := 0; i < r.field.NumX; i++ {
		for j := 0; j < r.field.NumY; j++ {
			if r.field.matrix[i][j] == true {
				r.DrawPoint(i, j)
			}
		}
	}
}

func (r *Game) DrawFigure() {
	for _, point := range r.figure.a {
		r.DrawPoint(point.x, point.y)
	}
}

func (r *Game) DrawScores() {
	text.Draw(r.screen, fmt.Sprintf("%d", r.field.cntDel), r.font, r.tx+r.field.width+130, r.ty+35, color.White)
}

func (r *Game) DrawPoint(x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x*CubeWidth+r.tx), float64(y*CubeWidth+r.ty))
	err := r.screen.DrawImage(r.Square.sprite, op)
	fail.Check(err)
}

func (r *Game) SetNewFigure() {
	r.figure = r.nextFig
	if err := r.figure.StartPosition(); err != nil {
		r.EndGame()
		return
	}
	r.nextFig = NewFig(r.field, RandomNum())
}

func (r *Game) DrawNextFigure() {
	x := r.tx + 545.0
	y := r.ty + 115.0
	if r.nextFig.Type.IsI() {
		x = r.tx + 520.0
		y = r.ty + 125.0
	}
	for _, point := range r.nextFig.a {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(point.x*CubeWidth+x), float64(point.y*CubeWidth+y))
		err := r.screen.DrawImage(r.Square.sprite, op)
		fail.Check(err)
	}
}

func (r *Game) EndGame() {
	r.isEnd = true
}

func (r *Game) GameOver() {
	r.field.Clear()
	r.DrawResults()
	if r.startEvent() {
		r.StartGame()
	}
}

func (r *Game) DrawResults() {
	text.Draw(r.screen, fmt.Sprintf("%d", r.field.cntDel), r.font, r.tx+r.field.width+130, r.ty+35, color.White)
	text.Draw(r.screen, fmt.Sprintf(
		"Your score: %d\n\nTo start new game\npress Enter",
		r.field.cntDel), r.font, r.tx+r.field.width/3, r.ty+r.field.width/4, color.White)
}

func (r Game) startEvent() bool {
	return ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter)
}

func (r *Game) StartGame() {
	r.isEnd = false
	r.field.cntDel = 0
	r.SetNewFigure()
}

func RandomNum() Tetromino {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Tetromino(generator.Intn(7))
}

func (r *Game) listenMoving() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		r.tact(r.figure.MoveRight, SideDelta)
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		r.tact(r.figure.MoveLeft, SideDelta)
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		r.figure.Rotate()
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		r.SetDelta(SmallDelta)
	case inpututil.IsKeyJustReleased(ebiten.KeyDown):
		r.SetDelta(Delta)
	}
}

func (r *Game) SetDelta(delta float64) {
	if r.delta != delta {
		r.delta = delta
	}
}

func (r *Game) FallDown() {
	if r.figure.IsStopped() {
		r.SetDelta(Delta)
		r.SetNewFigure()
		return
	}

	if r.tick {
		r.figure.FallDown(&r.field)
	}
}

type Point struct {
	x, y int
}

func tick() float64 {
	return float64(time.Now().UnixNano() / 1e6)
}

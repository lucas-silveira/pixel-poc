package basic

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lucas-silveira/pixel-poc/config"
	"github.com/lucas-silveira/pixel-poc/pkg/graph"
	"golang.org/x/image/colornames"
)

var walls []graph.Line

type Player struct {
	pos pixel.Vec
	*imdraw.IMDraw
}

func (p *Player) move(v pixel.Vec) {
	p.pos = p.pos.Sub(v)
	p.Reset()
	p.Clear()
	p.Color = colornames.Deeppink
	p.Push(p.pos)
	p.Circle(10, 0)
}

func (p *Player) handleMovement(win *pixelgl.Window, dt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		p.move(pixel.V(200*dt, 0))
	}
	if win.Pressed(pixelgl.KeyRight) {
		p.move(pixel.V(-200*dt, 0))
	}
	if win.Pressed(pixelgl.KeyDown) {
		p.move(pixel.V(0, 200*dt))
	}
	if win.Pressed(pixelgl.KeyUp) {
		p.move(pixel.V(0, -200*dt))
	}
}

func makePlayer(v pixel.Vec) *Player {
	imd := imdraw.New(nil)
	imd.Color = colornames.Deeppink
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(v)
	imd.Circle(10, 0)

	return &Player{v, imd}
}

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Basic graphs",
		Bounds: pixel.R(0, 0, config.ScreenWidth, config.ScreenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	player := makePlayer(win.Bounds().Center())

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		player.handleMovement(win, dt)

		win.Clear(color.Black)

		for _, l := range walls {
			l.Draw(win)
		}
		player.Draw(win)

		win.Update()
	}
}

package basic

import (
	"image/color"
	"math"
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

func (p *Player) newPos(v pixel.Vec) pixel.Vec {
	return p.pos.Sub(v)
}

func (p *Player) move(v pixel.Vec) {
	p.pos = p.newPos(v)
	p.Reset()
	p.Clear()
	p.Color = colornames.Deeppink
	p.Push(p.pos)
	p.Circle(10, 0)
}

func makePlayer(v pixel.Vec) *Player {
	imd := imdraw.New(nil)
	imd.Color = colornames.Deeppink
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(v)
	imd.Circle(10, 0)

	return &Player{v, imd}
}

func intersect(v pixel.Vec, l graph.Line) bool {
	line := l.Props

	m := (line.B.Y - line.A.Y) / (line.B.X - line.A.X) // Angular coefficient
	isVertical := math.IsInf(m, 0)

	if isVertical {
		dt := math.Abs(v.X - line.A.X)
		return dt <= 10
	}

	n := line.B.Y - m*line.B.X        // Linear coefficient
	y := m*v.X + n                    // Reduced equation of the line: y = mx + n
	return y >= v.Y-10 && y <= v.Y+10 // is collinear
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

	win.SetSmooth(true)

	walls = append(
		walls,
		*graph.NewLine(pixel.V(500, 110), pixel.V(200, 150), config.WallColor),
		*graph.NewLine(pixel.V(200, 400), pixel.V(100, 300), config.WallColor),
		*graph.NewLine(pixel.V(550, 400), pixel.V(680, 300), config.WallColor),
	)
	walls = append(walls, graph.NewLineRect(
		config.Padding,
		config.Padding,
		config.ScreenWidth-2*config.Padding,
		config.ScreenHeight-2*config.Padding,
		config.WallColor,
	)...)
	player := makePlayer(win.Bounds().Center())

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if win.Pressed(pixelgl.KeyLeft) {
			v := pixel.V(200*dt, 0)
			newPos := player.newPos(v)

			canMove := true
			for _, l := range walls {
				if intersect(newPos, l) {
					canMove = false
					break
				}
			}
			if canMove {
				player.move(v)
			}
		}
		if win.Pressed(pixelgl.KeyRight) {
			v := pixel.V(-200*dt, 0)
			newPos := player.newPos(v)

			canMove := true
			for _, l := range walls {
				if intersect(newPos, l) {
					canMove = false
					break
				}
			}
			if canMove {
				player.move(v)
			}
		}
		if win.Pressed(pixelgl.KeyDown) {
			v := pixel.V(0, 200*dt)
			newPos := player.newPos(v)

			canMove := true
			for _, l := range walls {
				if intersect(newPos, l) {
					canMove = false
					break
				}
			}
			if canMove {
				player.move(v)
			}
		}
		if win.Pressed(pixelgl.KeyUp) {
			v := pixel.V(0, -200*dt)
			newPos := player.newPos(v)

			canMove := true
			for _, l := range walls {
				if intersect(newPos, l) {
					canMove = false
					break
				}
			}
			if canMove {
				player.move(v)
			}
		}

		win.Clear(color.Black)

		for _, l := range walls {
			l.Draw(win)
		}
		player.Draw(win)

		win.Update()
	}
}

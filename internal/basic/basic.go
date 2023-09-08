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

func intersection(v pixel.Vec, l graph.Line) (float64, float64, bool) {
	line := l.Props
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
	denom := (v.X-v.X+5)*(line.A.Y-line.B.Y) - (v.Y-v.Y+5)*(line.A.X-line.B.X)
	tNum := (v.X-line.A.X)*(line.A.Y-line.B.Y) - (v.Y-line.A.Y)*(line.A.X-line.B.X)
	uNum := (v.X-line.A.X)*(v.Y-v.Y+5) - (v.Y-line.A.Y)*(v.X-v.X+5)

	if denom == 0 {
		return 0, 0, false
	}

	t := tNum / denom
	if t > 1 || t < 0 {
		return 0, 0, false
	}

	u := uNum / denom
	if u > 1 || u < 0 {
		return 0, 0, false
	}

	x := v.X + t*(v.X+5-v.X)
	y := v.Y + t*(v.Y+5-v.Y)
	return x, y, true

}
func intersect(v pixel.Vec, l graph.Line) bool {
	_, _, ok := intersection(v, l)
	return ok
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
		*graph.NewLine(pixel.V(200, 150), pixel.V(500, 110), config.WallColor),
		*graph.NewLine(pixel.V(100, 300), pixel.V(200, 400), config.WallColor),
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

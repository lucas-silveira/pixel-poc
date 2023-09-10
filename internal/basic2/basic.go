package basic2

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

func (p *Player) handleMovement(win *pixelgl.Window, dt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		v := pixel.V(200*dt, 0)
		newPos := p.newPos(v)

		canMove := true
		for _, l := range walls {
			if intersect(newPos, l) {
				canMove = false
				break
			}
		}
		if canMove {
			p.move(v)
		}
	}
	if win.Pressed(pixelgl.KeyRight) {
		v := pixel.V(-200*dt, 0)
		newPos := p.newPos(v)

		canMove := true
		for _, l := range walls {
			if intersect(newPos, l) {
				canMove = false
				break
			}
		}
		if canMove {
			p.move(v)
		}
	}
	if win.Pressed(pixelgl.KeyDown) {
		v := pixel.V(0, 200*dt)
		newPos := p.newPos(v)

		canMove := true
		for _, l := range walls {
			if intersect(newPos, l) {
				canMove = false
				break
			}
		}
		if canMove {
			p.move(v)
		}
	}
	if win.Pressed(pixelgl.KeyUp) {
		v := pixel.V(0, -200*dt)
		newPos := p.newPos(v)

		canMove := true
		for _, l := range walls {
			if intersect(newPos, l) {
				canMove = false
				break
			}
		}
		if canMove {
			p.move(v)
		}
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

func intersection(l1, l2 graph.Line) (float64, float64, bool) {
	line1 := l1.Props
	line2 := l2.Props
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
	denom := (line1.A.X-line1.B.X)*(line2.A.Y-line2.B.Y) - (line1.A.Y-line1.B.Y)*(line2.A.X-line2.B.X)
	tNum := (line1.A.X-line2.A.X)*(line2.A.Y-line2.B.Y) - (line1.A.Y-line2.A.Y)*(line2.A.X-line2.B.X)
	uNum := (line1.A.X-line2.A.X)*(line1.A.Y-line1.B.Y) - (line1.A.Y-line2.A.Y)*(line1.A.X-line1.B.X)

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

	x := line1.A.X + t*(line1.B.X-line1.A.X)
	y := line1.A.Y + t*(line1.B.Y-line1.A.Y)
	return x, y, true

}

func intersect(v pixel.Vec, l2 graph.Line) bool {
	l1 := graph.NewLine(v, pixel.V(v.X+5, v.Y+5), 1, color.White)
	_, _, ok := intersection(*l1, l2)
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

	walls = append(
		walls,
		*graph.NewLineByAngle(pixel.V(150, 170), 150, 1.565, config.LStroke, config.WallColor),
		*graph.NewLineByAngle(pixel.V(600, 170), 150, 1.565, config.LStroke, config.WallColor),
	)
	walls = append(walls, graph.NewLineRect(
		config.Padding,
		config.Padding,
		config.ScreenWidth-2*config.Padding,
		config.ScreenHeight-2*config.Padding,
		config.LStroke,
		config.WallColor,
	)...)
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

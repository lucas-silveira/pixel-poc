package basic

import (
	"image/color"
	"math"
	"sort"
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

// rayCasting returns a slice of line originating from point cx, cy and intersecting with objects
func rayCasting(center pixel.Vec, lines []graph.Line) []*graph.Line {
	const rayLength = 1000 // something large enough to reach all objects

	var rays []*graph.Line
	for _, l := range lines {
		// Cast two rays per point
		for _, p := range l.Points() {
			ray := graph.NewLine(center, p, 1, colornames.Salmon)
			angle := ray.Angle()

			for _, offset := range []float64{-0.0005, 0.0005} {
				points := []pixel.Vec{}
				ray2 := graph.NewLineByAngle(center, rayLength, angle+offset, 0, colornames.White)

				// Iterate over all lines to find the intersection points
				for _, l2 := range lines {
					if px, py, ok := intersection(*ray2, l2); ok {
						// Append the intersection point
						points = append(points, pixel.V(px, py))
					}
				}

				// Find the point closest to start of ray
				min := math.Inf(1)
				minIdx := -1
				for i, p2 := range points {
					d := math.Abs(center.X-p2.X) + math.Abs(center.Y-p2.Y) // distance between two points
					if d < min {
						min = d
						minIdx = i
					}
				}

				// Append the ray as well as a ray with the end point with the closest intersections
				rays = append(rays, ray, graph.NewLine(center, points[minIdx], 1, colornames.Blue))
			}
		}
	}

	sort.Slice(rays, func(i, j int) bool {
		return rays[i].Angle() < rays[j].Angle()
	})
	return rays
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
		rays := rayCasting(player.pos, walls)

		win.Clear(color.Black)

		for _, l := range walls {
			l.Draw(win)
		}
		for i, r := range rays {
			nextLine := rays[(i+1)%len(rays)]                   // if the next index is out of range, get back to the first one.
			color := pixel.RGB(1, 1, 1).Mul((pixel.Alpha(0.1))) // transparent color
			t := graph.NewTriangle(player.pos, r.Props.B, nextLine.Props.B, color)

			t.Draw(win)
			// r.Draw(win) // Draw lines
		}
		player.Draw(win)

		win.Update()
	}
}

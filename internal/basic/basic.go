package basic

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lucas-silveira/pixel-poc/config"
	"github.com/lucas-silveira/pixel-poc/pkg/graph"
)

var walls []graph.Line

func makePlayer(v pixel.Vec) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color.RGBA{255, 99, 71, 255}
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(v)
	imd.Circle(10, 0)

	return imd
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

	outsideWalls := graph.NewLineRect(
		config.Padding,
		config.Padding,
		config.ScreenWidth-2*config.Padding,
		config.ScreenHeight-2*config.Padding,
		config.WallColor,
	)
	walls = append(
		walls,
		*graph.NewLine(pixel.V(500, 110), pixel.V(200, 150), config.WallColor),
		*graph.NewLine(pixel.V(200, 400), pixel.V(100, 300), config.WallColor),
		*graph.NewLine(pixel.V(550, 400), pixel.V(680, 300), config.WallColor),
	)
	player := makePlayer(win.Bounds().Center())

	line := graph.NewLineByAngle(win.Bounds().Center(), 100, 0, config.WallColor)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		line.IncrementAngle(-1 * dt)

		win.Clear(color.Black)

		line.Draw(win)

		for _, l := range outsideWalls {
			l.Draw(win)
		}
		for _, l := range walls {
			l.Draw(win)
		}
		player.Draw(win)

		win.Update()
	}
}

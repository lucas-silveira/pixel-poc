package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	screenWidth  = 768
	screenHeight = 480
	lStroke      = 5
)

var walls []imdraw.IMDraw

type Line struct {
	X1, Y1, X2, Y2 float64
}

func (l *Line) angle() float64 {
	return math.Atan2(l.Y2-l.Y1, l.X2-l.X1)
}

func rect(x, y, w, h float64) []Line {
	return []Line{
		{x, y, x, y + h},
		{x, y + h, x + w, y + h},
		{x + w, y + h, x + w, y},
		{x + w, y, x, y},
	}
}

func makeExternalWall(c color.Color) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	imd.EndShape = imdraw.RoundEndShape
	var padding float64 = lStroke / 2
	lines := rect(padding, padding, screenWidth-2*padding, screenHeight-2*padding)

	for _, l := range lines {
		imd.Push(pixel.V(l.X1, l.Y1))
	}

	imd.Push(pixel.V(lines[3].X2, lines[3].Y2))
	imd.Line(lStroke)

	return imd
}

func makeInternalWall(c color.Color, l *Line) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = c
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(l.X1, l.Y1), pixel.V(l.X2, l.Y2))
	imd.Line(lStroke)

	return imd
}

func makePlayer(v pixel.Vec) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = color.RGBA{255, 99, 71, 255}
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(v)
	imd.Circle(10, 0)

	return imd
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	outsideWall := makeExternalWall(pixel.RGB(1, 1, 1))
	walls = append(
		walls,
		*makeInternalWall(pixel.RGB(1, 1, 1), &Line{500, 110, 200, 150}),
		*makeInternalWall(pixel.RGB(1, 1, 1), &Line{200, 400, 100, 300}),
		*makeInternalWall(pixel.RGB(1, 1, 1), &Line{550, 400, 680, 300}),
	)
	player := makePlayer(win.Bounds().Center())

	for !win.Closed() {
		win.Clear(color.Black)
		outsideWall.Draw(win)
		for _, l := range walls {
			l.Draw(win)
		}
		player.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

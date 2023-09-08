package sinfunc

import (
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lucas-silveira/pixel-poc/config"
)

type Wave struct {
	*imdraw.IMDraw
}

func (w *Wave) changeHight(h float64) {
	w.Reset()
	w.Clear()
	var x float64
	for x = 0; x < 768; x++ {
		/**
		Fórmula:
		f(x) = sin(x)
		g(x) = a*f(b*(x - c)) + d
		*/
		y := h*math.Sin(0.3*(x-1)) + config.ScreenHeight/2
		w.Push(pixel.V(x*10, y))
	}
	w.Line(1)
}

func newWave(c color.Color) *Wave {
	imd := imdraw.New(nil)
	imd.Color = c
	imd.EndShape = imdraw.RoundEndShape
	var x float64
	for x = 0; x < 768; x++ {
		/**
		Fórmula:
		f(x) = sin(x)
		g(x) = a*f(b*(x - c)) + d
		*/
		y := 100*math.Sin(0.3*(x-1)) + config.ScreenHeight/2
		imd.Push(pixel.V(x*10, y))
	}
	imd.Line(1)

	return &Wave{imd}
}

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sin Function",
		Bounds: pixel.R(0, 0, config.ScreenWidth, config.ScreenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	wave := newWave(pixel.RGB(1, 1, 1))

	var (
		minH, maxH    float64 = -100, 100
		h, speed, dir float64 = float64(minH), 150, 1
		last                  = time.Now()
	)
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if (dir == 1 && h >= maxH) || (dir == -1 && h <= minH) {
			dir *= -1
		}
		h += speed * dt * dir

		wave.changeHight(h)

		win.Clear(color.Black)

		wave.Draw(win)

		win.Update()
	}
}

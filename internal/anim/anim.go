package anim

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lucas-silveira/pixel-poc/config"
)

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

	imd := imdraw.New(nil)
	imd.Color = color.RGBA{255, 99, 71, 255}
	imd.EndShape = imdraw.RoundEndShape

	var (
		t, x, dir          float64   = 0, 100, 1
		minWidth, maxWidth float64   = 100, config.ScreenWidth - 100
		last               time.Time = time.Now()
	)

	for i := 0; i < 1000; i++ {
		fmt.Println(math.Sin(float64(i)))
	}

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		t += dt
		last = time.Now()

		if dir == 1 && x >= maxWidth {
			dir = -1
		}
		if dir == -1 && x <= minWidth {
			dir = 1
		}

		x = config.ScreenWidth/2 + ((config.ScreenWidth / 2) * math.Sin(t))

		imd.Reset()
		imd.Clear()
		imd.Push(pixel.V(x, config.ScreenHeight/2))
		imd.Circle(10, 0)

		win.Clear(color.Black)

		imd.Draw(win)

		win.Update()
	}
}

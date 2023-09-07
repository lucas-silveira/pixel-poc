package sinfunc

import (
	"fmt"
	"image/color"
	"math"

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

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)
	var x float64
	imd.EndShape = imdraw.RoundEndShape
	for x = 0; x < 768; x++ {
		y := math.Sin(x)*10 + config.ScreenHeight/2
		fmt.Println(y)
		imd.Push(pixel.V(x*10, y))
	}
	imd.Line(1)

	for !win.Closed() {

		win.Clear(color.Black)
		imd.Draw(win)

		win.Update()
	}
}

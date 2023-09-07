package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/lucas-silveira/pixel-poc/internal/anim"
)

func main() {
	pixelgl.Run(anim.Run)
}

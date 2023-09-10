package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/lucas-silveira/pixel-poc/internal/raycasting"
)

func main() {
	pixelgl.Run(raycasting.Run)
}

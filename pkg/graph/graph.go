package graph

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/lucas-silveira/pixel-poc/config"
)

// Line is a single dimension graph
type Line struct {
	Props pixel.Line
	imdraw.IMDraw
}

// Angle returns the angle of the line
func (l *Line) Angle() float64 {
	return math.Atan2(l.Props.B.Y-l.Props.A.Y, l.Props.B.X-l.Props.A.X)
}

// Length returns the lenght of the line
func (l *Line) Length() float64 {
	xDelta := l.Props.B.X - l.Props.A.X
	yDelta := l.Props.B.Y - l.Props.A.Y
	// Pythagorean theorem
	return math.Sqrt(math.Pow(xDelta, 2) + math.Pow(yDelta, 2))
}

// IncrementAngle increment the line angle by i
func (l *Line) IncrementAngle(i float64) {
	xDelta := l.Length() * math.Cos(l.Angle()+i) // Ax = A*cos(θ)
	yDelta := l.Length() * math.Sin(l.Angle()+i) // Ay = A*sin(θ)
	x2 := l.Props.A.X + xDelta
	y2 := l.Props.A.Y + yDelta

	l.Props.B.X = x2
	l.Props.B.Y = y2

	l.Reset()
	l.Clear()
	l.Push(l.Props.A, l.Props.B)
	l.Line(config.LStroke)
}

// NewLine creates a line and return it
func NewLine(begin, end pixel.Vec, c color.Color) *Line {
	l := &Line{
		pixel.L(
			begin,
			end,
		),
		*imdraw.New(nil),
	}
	l.Color = c
	l.EndShape = imdraw.RoundEndShape
	l.Push(l.Props.A, l.Props.B)
	l.Line(config.LStroke)

	return l
}

// NewLineByAngle creates a line by an angle and return it
func NewLineByAngle(begin pixel.Vec, length, angle float64, c color.Color) *Line {
	xDelta := length * math.Cos(angle) // Componente Ax = A*cos(θ)
	yDelta := length * math.Sin(angle) // Componente Ay = A*sin(θ)
	x2 := begin.X + xDelta
	y2 := begin.Y + yDelta

	l := &Line{
		pixel.L(
			begin,
			pixel.V(x2, y2),
		),
		*imdraw.New(nil),
	}
	l.Color = c
	l.EndShape = imdraw.RoundEndShape
	l.Push(l.Props.A, l.Props.B)
	l.Line(config.LStroke)

	return l
}

// NewLineRect creates a rect line and return it
func NewLineRect(x, y, w, h float64, c color.Color) []Line {
	return []Line{
		*NewLine(pixel.V(x, y), pixel.V(x, y+h), c),
		*NewLine(pixel.V(x, y+h), pixel.V(x+w, y+h), c),
		*NewLine(pixel.V(x+w, y+h), pixel.V(x+w, y), c),
		*NewLine(pixel.V(x+w, y), pixel.V(x, y), c),
	}
}

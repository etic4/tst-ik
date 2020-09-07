package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	vec "github.com/machinbrol/vecmaths"
)

type line struct {
	a      vec.Vec2
	b      vec.Vec2
	angle  float64
	length float64
}

func newLine(a vec.Vec2, angle float64, length float64) *line {
	l := new(line)
	l.a = a
	l.angle = angle
	l.length = length
	l.setB()
	return l
}

//définit la position de b
func (l *line) setB() {
	bPos := vec.FromPolar(l.angle, l.length)
	l.b = l.a.Add(bPos)
}

func (l *line) setAngle(angle float64) {
	l.angle = angle
	l.setB()
}

//rotation en radians autour du point a
func (l *line) rotate(angle float64) {
	l.angle += angle
	l.setB()
}

func (l *line) draw(thick int, clr rl.Color) {
	a := rl.Vector2{X: float32(l.a.X), Y: float32(l.a.Y)}
	b := rl.Vector2{X: float32(l.b.X), Y: float32(l.b.Y)}
	rl.DrawLineEx(a, b, float32(thick), clr)
}

func (l *line) String() string {
	return fmt.Sprintf("%v-%v", l.a, l.b)
}

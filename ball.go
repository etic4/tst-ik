package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	vec "github.com/machinbrol/vecmaths"
)

type ball struct {
	center vec.Vec2
	radius float64
	speed  vec.Vec2
	color  rl.Color
}

func newBall(center vec.Vec2, radius float64, color rl.Color) *ball {
	b := new(ball)
	b.center = center
	b.radius = radius
	b.color = color

	return b
}

func (b *ball) getSpeed() vec.Vec2 {
	return b.speed
}

func (b *ball) setSpeed(speed vec.Vec2) {
	b.speed = speed
}

func (b *ball) move() {
	b.center = b.center.Add(b.speed)
}

func (b *ball) collideLeft(x float64) bool {
	return b.center.X-b.radius < x
}

func (b *ball) collideRight(x float64) bool {
	return b.center.X+b.radius > x
}

func (b *ball) collideTop(y float64) bool {
	return b.center.Y-b.radius < y
}

func (b *ball) collideBottom(y float64) bool {
	return b.center.Y+b.radius > y
}

func (b *ball) draw() {
	//rl.DrawCircle(int32(b.center.X), int32(b.center.Y), float32(b.radius), clr)
	rl.DrawCircleLines(int32(b.center.X), int32(b.center.Y), float32(b.radius), b.color)
}

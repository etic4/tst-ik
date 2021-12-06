package main

import vec "github.com/machinbrol/vecmath"

type mover interface {
	move()
	getSpeed() vec.Vec2
	setSpeed(vec.Vec2)
	collideLeft(float64) bool
	collideRight(float64) bool
	collideTop(float64) bool
	collideBottom(float64) bool
}

type scene struct {
	origin  vec.Vec2
	width   float64
	height  float64
	objects []mover
}

func newScene(origin vec.Vec2, width float64, height float64) *scene {
	s := new(scene)
	s.origin = origin
	s.width = width
	s.height = height
	return s
}

func (s *scene) add(obj mover) {
	s.objects = append(s.objects, obj)
}

func (s *scene) move(obj mover) {
	speed := obj.getSpeed()

	if obj.collideLeft(0) || obj.collideRight(s.width) {
		speed.X = -speed.X
		obj.setSpeed(speed)
	}
	if obj.collideTop(0) || obj.collideBottom(s.height) {
		speed.Y = -speed.Y
		obj.setSpeed(speed)
	}
	obj.move()
}

func (s *scene) update() {
	for _, obj := range s.objects {
		s.move(obj)
	}
}

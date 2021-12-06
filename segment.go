package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	vec "github.com/machinbrol/vecmath"
)

type segment struct {
	*line
	parent *segment
	child  *segment
	thick  int
	color  rl.Color
	name   string
}

func newSegment(a vec.Vec2, angle float64, length float64) *segment {
	s := new(segment)
	s.init(a, angle, length)
	return s
}

func newChild(parent *segment, angle float64, length float64) *segment {
	s := new(segment)
	s.init(parent.b, angle, length)
	s.parent = parent
	s.parent.setChild(s)
	return s
}

func (s *segment) init(a vec.Vec2, angle float64, length float64) {
	s.line = new(line)
	s.a = a
	s.angle = angle
	s.length = length
	s.setB()
	s.thick = 1
	if s.color == *new(rl.Color) {
		s.color = White
	}
}

func (s *segment) setChild(child *segment) {
	s.child = child
}

func (s *segment) rotate(angle float64) {
	if s.parent != nil {
		s.a = s.parent.b
		s.setB()
	}

	s.line.rotate(angle)

	if s.child != nil {
		s.child.rotate(angle)
	}
}

func (s *segment) getAngleAndVecTo(pos vec.Vec2) (float64, vec.Vec2) {
	//l'angle vers pos
	direction := pos.Sub(s.a)
	angle := direction.Angle()

	//la position à atteindre pour 'a' pour que 'b' soit sur pos
	//le vecteur direction est mis à la longeur
	//du segment et retourné puis l'origine est positionnée sur la cible
	targetPos := direction.SetMag(s.length)
	targetPos = targetPos.Mult(-1)
	targetPos = targetPos.Add(pos)

	return angle, targetPos
}

func (s *segment) moveHeadTo(pos vec.Vec2) {
	angle, targetPos := s.getAngleAndVecTo(pos)

	s.setAngle(angle)
	s.a = targetPos
	s.setB()

	if s.parent != nil {
		s.parent.moveHeadTo(s.a)
	}
}

func (s *segment) moveBack(pos vec.Vec2) {
	s.a = pos
	s.setB()
	if s.child != nil {
		s.child.moveBack(s.b)
	}
}

func (s *segment) getHead() *segment {
	if s.child != nil {
		return s.child.getHead()
	}
	return s
}

func (s *segment) draw() {
	s.line.draw(s.thick, s.color)
	if s.child != nil {
		s.child.draw()
	}
}

func (s *segment) String() string {
	return fmt.Sprintf("segment %v", s.line)
}

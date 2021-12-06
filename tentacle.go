package main

import (
	vec "github.com/machinbrol/vecmath"
	"github.com/machinbrol/vecmath/maths"
)

var (
	defaulAngle     float64 = maths.Rad(90)
	defaultMaxForce float64 = 3
)

type tentacle struct {
	base         vec.Vec2
	nbrSegs      int
	segSize      float64
	speed        vec.Vec2
	maxSpeed     float64
	maxForce     float64
	detectionRay float64
	angle        float64
	thick        int
	restPos      vec.Vec2
	segments     []*segment
	*segment
}

func newTentacle(base vec.Vec2, nbrSegs int, segSize float64, speed vec.Vec2, maxSpeed float64, maxForce float64, thick int, detectionRay float64) *tentacle {
	t := new(tentacle)
	t.base = base
	t.nbrSegs = nbrSegs
	t.segSize = segSize
	t.speed = speed
	t.maxSpeed = maxSpeed
	t.maxForce = maxForce
	t.thick = thick
	t.detectionRay = detectionRay

	t.angle = defaulAngle

	t.build()
	return t
}

func (t *tentacle) build() {
	t.segment = newSegment(t.base, t.angle, t.segSize)
	t.segment.thick = t.thick

	parent := t.segment
	for i := 1; i < t.nbrSegs; i++ {
		next := newChild(parent, parent.angle, parent.length)
		next.thick = int(maths.Lerp(float64(t.thick), float64(1), float64(i)/float64(t.nbrSegs)))
		parent = next
	}
	t.restPos = parent.b
}

//followNearest suit le plus près
func (t *tentacle) followNearest(objects []*ball) {
	nearest := t.getNearest(objects)
	t.followSteering(nearest.center)
}

func (t *tentacle) followSteering(pos vec.Vec2) {
	if pos == (vec.Vec2{}) || !t.isInDetectionRay(pos) {
		pos = t.restPos
	}
	desired := pos.Sub(t.getHead().b)
	desired = desired.LimitMag(-t.maxSpeed, t.maxSpeed)
	steer := desired.Sub(t.speed)
	steer = steer.LimitMag(-t.maxForce, t.maxForce)
	t.speed = t.speed.Add(steer)
	t.followPos(t.getHead().b.Add(t.speed))
}

//followPos suit une position
func (t *tentacle) followPos(pos vec.Vec2) {
	base := t.a
	t.getHead().moveHeadTo(pos)
	t.moveBack(base)
}

func (t *tentacle) isInDetectionRay(pos vec.Vec2) bool {
	return t.getHead().b.Distance(pos) <= t.detectionRay
}

//getNearest retourne le plus proche
func (t *tentacle) getNearest(objects []*ball) *ball {
	detected := []*ball{}

	for _, obj := range objects {
		if t.isInDetectionRay(obj.center) {
			detected = append(detected, obj)
		}
	}

	var nearest = new(ball)

	if len(detected) > 0 {
		nearest = detected[0]

		for i := 1; i < len(detected); i++ {
			if t.b.Distance(detected[i].center) < t.b.Distance(nearest.center) {
				nearest = detected[i]
			}
		}
	}
	return nearest
}

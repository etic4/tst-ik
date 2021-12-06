// Cf. The Coding Train https://www.youtube.com/watch?v=hbgDqyy8bIw

package main

import (
	"math/rand"

	vec "github.com/etic4/vecmath"
)

var (
	screenWidth      float64 = 800
	screenHeight     float64 = 600
	fps                      = 60
	tentac0          *tentacle
	tentac1          *tentacle
	tentacles        []*tentacle = []*tentacle{}
	nbrTentacles     int         = 1
	betweenTentacles float64     = 50
	balles           []*ball     = []*ball{}
	scene0           *scene
	segSize          float64  = 8.0
	nbrSegs          int      = 35
	thick            int      = 5
	speed            vec.Vec2 = vec.ZERO()
	maxSpeed         float64  = 2.0
	maxForce         float64  = 1.0
	detectionRay     float64  = 200.0
	nbrBalles        int      = 3
	ballRay          float64  = 10.0
)

func init() {
	scene0 = newScene(vec.Vec2{X: 0, Y: 0}, screenWidth, screenHeight)

	getTentacles(nbrTentacles)
	getBalls(nbrBalles)

}

func getBalls(nbr int) {
	//création balle(s)
	for i := 0; i < nbrBalles; i++ {
		b := newBall(vec.Vec2{X: float64(100 + rand.Intn(100)), Y: float64(100 + rand.Intn(100))}, ballRay, White)
		b.setSpeed(vec.Vec2{X: float64(2 + rand.Intn(3)), Y: float64(2 + rand.Intn(3))})
		balles = append(balles, b)
		scene0.add(b)
	}
}

func getTentacles(nbr int) {
	between := betweenTentacles
	start := screenWidth/2 - (float64(nbr)*between)/2

	for i := 0; i < nbr; i++ {
		start += between/2 + float64(rand.Intn(int(between/2)))
		pos := vec.NewVector(start, screenHeight-200+float64(rand.Intn(100)))
		segs := nbrSegs - rand.Intn(5)
		size := segSize - float64(rand.Intn(3))
		maxSpeed := maxSpeed * float64(100-50+rand.Intn(100)) / 100
		maxForce := maxForce * float64(100-50+rand.Intn(100)) / 100
		thick := thick + (rand.Intn(6) - 3)
		detectionRay := detectionRay * float64(100-50+rand.Intn(100)) / 100
		tent := newTentacle(pos, segs, size, speed, maxSpeed, maxForce, thick, detectionRay)
		tentacles = append(tentacles, tent)
	}
}

func drawBalls() {
	for _, b := range balles {
		b.draw()
	}
}

func drawTentacles() {
	for _, tent := range tentacles {
		tent.draw()
	}
}

func update() {
	//tentac0.followSteering(mousPos())
	for _, tent := range tentacles {
		tent.followNearest(balles)
	}
	scene0.update()
}

func draw() {
	drawTentacles()
	drawBalls()
}

func main() {
	run(screenWidth, screenHeight, fps)
}

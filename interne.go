package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	vec "github.com/machinbrol/vecmath"
)

var (
	//White couleur blanc
	White = rl.White

	//Black couleur noir
	Black = rl.Black
)

type raylibOptions struct {
	fps          int
	screenWidth  int
	screenHeight int
	winTitle     string
}

func mouseX() float64 {
	return float64(rl.GetMouseX())
}

func mouseY() float64 {
	return float64(rl.GetMouseY())
}

func mousPos() vec.Vec2 {
	return vec.NewVector(mouseX(), mouseY())
}

func run(screenWidth float64, screenHeight float64, fps int) {
	//Fps := int32(fps)

	/* Synchronisation verticale
	Si je comprend bien, ça aligne le fps sur le taux de rafaichissement du moniteur*/
	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(int32(screenWidth), int32(screenHeight), "raylib")

	//rl.SetTargetFPS(Fps)  //nécessaire si synchronisation verticale ?

	for !rl.WindowShouldClose() {
		update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

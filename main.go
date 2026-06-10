package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1366, 768, "OSMV")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Color{R: 29, G: 29, B: 32, A: 255})
		rl.DrawText("Projetistas Software 3000", 800, 600, 30, rl.White)

		rl.EndDrawing()
	}
}

package main

import (
	"github.com/fayrghos/osmv/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1366, 768, "OSMV")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	fonteSans := rl.LoadFont("./assets/fonts/LiberationSans.ttf")

	box1 := ui.Digibox{
		Pos:      rl.Vector2{X: 200, Y: 200},
		Tam:      rl.Vector2{X: 300, Y: 50},
		Texto:    "ABC",
		Fonte:    &fonteSans,
		FonteTam: 32,
	}

	box2 := ui.Digibox{
		Pos:      rl.Vector2{X: 200, Y: 400},
		Tam:      rl.Vector2{X: 300, Y: 50},
		Texto:    "ABC",
		Fonte:    &fonteSans,
		FonteTam: 32,
	}

	for !rl.WindowShouldClose() {
		box1.Atualizar()
		box2.Atualizar()

		rl.BeginDrawing()

		rl.ClearBackground(ui.CorFundo)
		rl.DrawText("Projetistas Software 3000", 800, 600, 30, rl.White)

		box1.Desenhar()
		box2.Desenhar()

		rl.EndDrawing()
	}
}

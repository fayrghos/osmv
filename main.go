package main

import (
	"github.com/fayrghos/osmv/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(1366, 768, "OSMV")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	fonteSans := rl.LoadFontEx("./assets/fonts/LiberationSans.ttf", 40, nil)

	box1 := ui.Digibox{
		Pos:      rl.Vector2{X: 200, Y: 200},
		Tam:      rl.Vector2{X: 300, Y: 50},
		Texto:    "1359",
		Fonte:    &fonteSans,
		FonteTam: 40,
	}

	for !rl.WindowShouldClose() {
		box1.Atualizar()

		rl.BeginDrawing()

		rl.ClearBackground(ui.CorFundo)
		rl.DrawText("Projetistas Software 3000", 800, 600, 30, rl.White)

		box1.Desenhar()

		rl.EndDrawing()
	}
}

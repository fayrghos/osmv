package main

import (
	"github.com/fayrghos/osmv/internal/ui"
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(1366, 768, "OSMV")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	fonteSans := utils.CarregarFonte("./assets/fonts/LiberationSans.ttf", 28)

	box1 := ui.Digibox{
		Pos: rl.Vector2{X: 200, Y: 200},
		Tam: rl.Vector2{X: 400, Y: 75},
		Campo: utils.Texto{
			Conteudo: "13967",
			Tam:      48,
			Fonte:    &fonteSans,
		},
		Titulo: utils.Texto{
			Conteudo: "Teste",
			Tam:      36,
			Fonte:    &fonteSans,
		},
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

package main

import (
	"github.com/fayrghos/osmv/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1366, 768, "OSMV")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	fonte := rl.LoadFont("assets/fontes/LiberationSans.ttf")

	botoes := make(ui.GrupoBotoes, 0, 10)
	botoes = append(botoes, ui.Botao{
		Pos:     rl.Vector2{X: 200, Y: 200},
		Tamanho: rl.Vector2{X: 200, Y: 50},
		Fonte:   &fonte,
	})

	botoes = append(botoes, ui.Botao{
		Pos:     rl.Vector2{X: 200, Y: 300},
		Tamanho: rl.Vector2{X: 200, Y: 50},
		Fonte:   &fonte,
	})

	for !rl.WindowShouldClose() {
		botoes.AtualizarTodos()

		rl.BeginDrawing()

		rl.ClearBackground(ui.CorFundo)
		rl.DrawText("Projetistas Software 3000", 800, 600, 30, rl.White)

		botoes.DesenharTodos()

		rl.EndDrawing()
	}
}

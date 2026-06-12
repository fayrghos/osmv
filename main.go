package main

import (
	"github.com/fayrghos/osmv/internal/screens"
	"github.com/fayrghos/osmv/internal/state"
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
	defer rl.UnloadFont(fonteSans)

	globais := state.Globais{
		FonteSans: &fonteSans,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(ui.CorFundo)

		switch globais.TelaAtual {
		case state.TelaInicial:
			screens.AtualizarInicial(&globais)
			screens.DesenharInicial(&globais)
		}

		rl.EndDrawing()
	}
}

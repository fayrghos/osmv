package main

import (
	"flag"

	"github.com/fayrghos/osmv/internal/screens"
	"github.com/fayrghos/osmv/internal/state"
	"github.com/fayrghos/osmv/internal/ui"
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	modoRapido := flag.Bool("rapido", false, "ativa o modo rápido")
	flag.Parse()

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(state.Larg, state.Altu, "OSMV")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetExitKey(rl.KeyNull)

	icone := rl.LoadImage("./assets/images/icon.png")
	defer rl.UnloadImage(icone)

	rl.SetWindowIcon(*icone)

	fonteSans := utils.CarregarFonte("./assets/fonts/LiberationSans.ttf", 28)
	defer rl.UnloadFont(fonteSans)

	globais := state.Globais{
		FonteSans:  &fonteSans,
		ModoRapido: *modoRapido,
		BoxErro: ui.Errobox{
			Pos: rl.Vector2{X: state.Larg/2 - 400, Y: 35},
			Tam: rl.Vector2{X: 800, Y: 60},
			Campo: utils.Texto{
				Tam:   32,
				Fonte: &fonteSans,
			},
		},
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(ui.CorFundo)

		switch globais.TelaAtual {
		case state.TelaInicial:
			screens.AtualizarInicial(&globais)
			screens.DesenharInicial(&globais)
		case state.TelaPrincipal:
			screens.AtualizarPrincipal(&globais)
			screens.DesenharPrincipal(&globais)
		}

		globais.BoxErro.Desenhar()
		rl.EndDrawing()
	}
}

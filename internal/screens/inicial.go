package screens

import (
	"strconv"

	"github.com/fayrghos/osmv/internal/state"
	"github.com/fayrghos/osmv/internal/ui"
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Redesenho da tela inicial
func DesenharInicial(globais *state.Globais) {
	globais.BoxTamPaginas.Desenhar()
	globais.BoxTamFisica.Desenhar()
	globais.BoxTamLogica.Desenhar()
}

// Redesenho da tela inicial
func AtualizarInicial(globais *state.Globais) {
	globais.InicializarTela(state.TelaInicial, func() {
		globais.BoxTamPaginas = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 - 250, Y: state.Altu/2 - 40 - 150},
			Tam: rl.Vector2{X: 500, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Tam. Páginas",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}
		globais.BoxTamFisica = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 - 250, Y: state.Altu/2 - 40},
			Tam: rl.Vector2{X: 500, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Tam. Memória Física",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}
		globais.BoxTamLogica = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 - 250, Y: state.Altu/2 - 40 + 150},
			Tam: rl.Vector2{X: 500, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Tam. Memória Lógica",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}
	})

	globais.BoxTamPaginas.Atualizar()
	globais.BoxTamFisica.Atualizar()
	globais.BoxTamLogica.Atualizar()

	if rl.IsKeyPressed(rl.KeyEnter) {
		var err error

		globais.TamPaginas, err = strconv.Atoi(globais.BoxTamPaginas.Campo.Conteudo)
		if err != nil {
			panic("O conteúdo do campo de páginas não pôde ser convertido")
		}

		globais.TamFisica, err = strconv.Atoi(globais.BoxTamFisica.Campo.Conteudo)
		if err != nil {
			panic("O conteúdo do campo da memória física não pôde ser convertido")
		}

		globais.TamLogica, err = strconv.Atoi(globais.BoxTamLogica.Campo.Conteudo)
		if err != nil {
			panic("O conteúdo do campo da memória lógica não pôde ser convertido")
		}

		globais.TelaAtual = state.TelaPrincipal
	}
}

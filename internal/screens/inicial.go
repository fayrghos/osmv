package screens

import (
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
		globais.TamPaginas = globais.BoxTamPaginas.Exportar()
		globais.TamFisica = globais.BoxTamFisica.Exportar()
		globais.TamLogica = globais.BoxTamLogica.Exportar()

		if globais.TamPaginas == 0 || globais.TamFisica == 0 || globais.TamLogica == 0 {
			globais.BoxErro.Definir("Todos os campos devem ser maiores que 0.")
		} else {
			globais.BoxErro.Definir("")
			globais.TelaAtual = state.TelaPrincipal
		}
	}
}

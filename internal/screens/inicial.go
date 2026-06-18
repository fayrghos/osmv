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
	globais.BoxNumBits.Desenhar()
	globais.BoxTamFisica.Desenhar()
	globais.BoxTamLogica.Desenhar()
	globais.BotaoContinuar.Desenhar()
}

// Lógica da tela inicial
func AtualizarInicial(globais *state.Globais) {
	globais.InicializarTela(state.TelaInicial, func() {
		globais.BoxTamPaginas = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 - 450, Y: state.Altu/2 - 40 - 100},
			Tam: rl.Vector2{X: 400, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Num. de Páginas",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}

		globais.BoxNumBits = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 - 450, Y: state.Altu/2 - 40 + 75},
			Tam: rl.Vector2{X: 400, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Num. Bits Processo",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}

		globais.BoxTamFisica = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 + 50, Y: state.Altu/2 - 40 - 100},
			Tam: rl.Vector2{X: 400, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Num. Bits Memória Física",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}

		globais.BoxTamLogica = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 + 50, Y: state.Altu/2 - 40 + 75},
			Tam: rl.Vector2{X: 400, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Num. de slots Memória Física",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
		}

		if globais.ModoRapido {
			globais.BoxTamPaginas.Campo.Conteudo = "4"
			globais.BoxTamFisica.Campo.Conteudo = "8"
			globais.BoxNumBits.Campo.Conteudo = "16"
			globais.BoxTamLogica.Campo.Conteudo = "4"
		}

		globais.BotaoContinuar = ui.Botao{
			Pos: rl.Vector2{X: state.Larg/2 - 100, Y: state.Altu - 200},
			Tam: rl.Vector2{X: 200, Y: 80},
			Rotulo: utils.Texto{
				Conteudo: "Avançar",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: false,
			Clique: func() {
				globais.TamPaginas = globais.BoxTamPaginas.Exportar()
				globais.NumBits = globais.BoxNumBits.Exportar()
				globais.TamFisica = globais.BoxTamFisica.Exportar()
				globais.TamLogica = globais.BoxTamLogica.Exportar()

				if globais.TamPaginas == 0 ||
					globais.NumBits == 0 ||
					globais.TamFisica == 0 ||
					globais.TamLogica == 0 {
					globais.BoxErro.Definir("Todos os campos devem ser maiores que 0.")
				} else {
					globais.BoxErro.Definir("")
					globais.TelaAtual = state.TelaPrincipal
				}
			},
		}
	})

	globais.BoxTamPaginas.Atualizar()
	globais.BoxNumBits.Atualizar()
	globais.BoxTamFisica.Atualizar()
	globais.BoxTamLogica.Atualizar()
	globais.BotaoContinuar.Atualizar()
}

package screens

import (
	"fmt"

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
				Conteudo: "Num. Páginas",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			ValorMax: 16,
			ValorMin: 2,
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
			ValorMax: 64,
			ValorMin: 16,
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
			ValorMax: 128,
			ValorMin: 8,
		}

		globais.BoxTamLogica = ui.Digibox{
			Pos: rl.Vector2{X: state.Larg/2 + 50, Y: state.Altu/2 - 40 + 75},
			Tam: rl.Vector2{X: 400, Y: 80},
			Campo: utils.Texto{
				Tam:   48,
				Fonte: globais.FonteSans,
			},
			Titulo: utils.Texto{
				Conteudo: "Num. Slots Memória Física",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			ValorMax: 16,
			ValorMin: 2,
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
				var err error

				globais.TamPaginas, err = globais.BoxTamPaginas.Exportar()
				if err != nil {
					globais.BoxErro.Definir(fmt.Sprintf(
						"O número de páginas deve estar entre %d e %d.",
						globais.BoxTamPaginas.ValorMin,
						globais.BoxTamPaginas.ValorMax,
					))
					return
				}

				globais.NumBits, err = globais.BoxNumBits.Exportar()
				if err != nil {
					globais.BoxErro.Definir(fmt.Sprintf(
						"Os bits dos processos devem estar entre %d e %d.",
						globais.BoxNumBits.ValorMin,
						globais.BoxNumBits.ValorMax,
					))
					return
				}

				globais.TamFisica, err = globais.BoxTamFisica.Exportar()
				if err != nil {
					globais.BoxErro.Definir(fmt.Sprintf(
						"Os bits da memória física devem estar entre %d e %d.",
						globais.BoxTamFisica.ValorMin,
						globais.BoxTamFisica.ValorMax,
					))
					return
				}

				globais.TamLogica, err = globais.BoxTamLogica.Exportar()
				if err != nil {
					globais.BoxErro.Definir(fmt.Sprintf(
						"Os slots da memória física devem estar entre %d e %d.",
						globais.BoxTamLogica.ValorMin,
						globais.BoxTamLogica.ValorMax,
					))
					return
				}

				globais.BoxErro.Definir("")
				globais.TelaAtual = state.TelaPrincipal
			},
		}
	})

	globais.BoxTamPaginas.Atualizar()
	globais.BoxNumBits.Atualizar()
	globais.BoxTamFisica.Atualizar()
	globais.BoxTamLogica.Atualizar()
	globais.BotaoContinuar.Atualizar()
}

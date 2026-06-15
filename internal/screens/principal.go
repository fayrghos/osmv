package screens

import (
	"fmt"

	"github.com/fayrghos/osmv/internal/state"
	"github.com/fayrghos/osmv/internal/ui"
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func AtualizarPrincipal(globais *state.Globais) {
	globais.InicializarTela(state.TelaPrincipal, func() {
		globais.BotaoContinuar = ui.Botao{
			Pos: rl.Vector2{X: 598, Y: 441},
			Tam: rl.Vector2{X: 268, Y: 75},
			Rotulo: utils.Texto{
				Conteudo: "Continuar",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: false,
			Clique: func() {
				fmt.Println("Clicaram em continuar!!")
			},
		}

		globais.BotaoPasso = ui.Botao{
			Pos: rl.Vector2{X: 598, Y: 528},
			Tam: rl.Vector2{X: 268, Y: 75},
			Rotulo: utils.Texto{
				Conteudo: "Passo",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: false,
			Clique: func() {
				fmt.Println("Clicaram em passo!!")
			},
		}

		globais.BotaoVoltar = ui.Botao{
			Pos: rl.Vector2{X: 598, Y: 616},
			Tam: rl.Vector2{X: 268, Y: 75},
			Rotulo: utils.Texto{
				Conteudo: "Voltar",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: true,
			Clique: func() {
				fmt.Println("Clicaram em voltar!!")
			},
		}
	})

	globais.BotaoContinuar.Atualizar()
	globais.BotaoPasso.Atualizar()
	globais.BotaoVoltar.Atualizar()
}

func DesenharPrincipal(globais *state.Globais) {
	// ------------------------------
	// Páginas
	// ------------------------------
	recPaginas := rl.Rectangle{
		X:      77,
		Y:      77,
		Width:  230,
		Height: 614,
	}
	rl.DrawRectangleRounded(recPaginas, 0.1, 4, ui.CorPrincipal)

	recPaginasIn := rl.Rectangle{
		X:      88,
		Y:      146,
		Width:  208,
		Height: 533,
	}
	rl.DrawRectangleRec(recPaginasIn, ui.CorSecundaria)

	// ------------------------------
	// Quadros
	// ------------------------------
	recQuadros := rl.Rectangle{
		X:      338,
		Y:      77,
		Width:  230,
		Height: 614,
	}
	rl.DrawRectangleRounded(recQuadros, 0.1, 4, ui.CorPrincipal)

	recrecQuadrosIn := rl.Rectangle{
		X:      348,
		Y:      146,
		Width:  208,
		Height: 533,
	}
	rl.DrawRectangleRec(recrecQuadrosIn, ui.CorSecundaria)

	// ------------------------------
	// Tabela
	// ------------------------------
	recTabela := rl.Rectangle{
		X:      598,
		Y:      77,
		Width:  268,
		Height: 352,
	}
	rl.DrawRectangleRounded(recTabela, 0.08, 4, ui.CorPrincipal)

	recrecTabelaIn := rl.Rectangle{
		X:      609,
		Y:      146,
		Width:  245,
		Height: 271,
	}
	rl.DrawRectangleRec(recrecTabelaIn, ui.CorSecundaria)

	// ------------------------------
	// Memória
	// ------------------------------
	recMemoria := rl.Rectangle{
		X:      896,
		Y:      77,
		Width:  393,
		Height: 614,
	}
	rl.DrawRectangleRounded(recMemoria, 0.05, 4, ui.CorPrincipal)

	recrecMemoriaIn := rl.Rectangle{
		X:      907,
		Y:      146,
		Width:  371,
		Height: 533,
	}
	rl.DrawRectangleRec(recrecMemoriaIn, ui.CorSecundaria)

	// ------------------------------
	// Botões
	// ------------------------------
	globais.BotaoContinuar.Desenhar()
	globais.BotaoPasso.Desenhar()
	globais.BotaoVoltar.Desenhar()

	type Tabela struct {
		PosPaginas rl.Vector2
		DimPaginas rl.Vector2
		PosQuadros rl.Vector2
		DimQuadros rl.Vector2
		Cor        rl.Color
		CorBorda   rl.Color
	}

	slotPaginas := alturaFixa / float32(globais.TamPaginas)

	for i := 1; i < globais.TamPaginas; i++ {
		posY := recPaginasIn.Y + float32(i)*slotPaginas
		posIni := rl.Vector2{X: recPaginasIn.X, Y: posY}
		posEnd := rl.Vector2{X: recPaginasIn.X + larguraFixa, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}

	for i := 1; i < globais.TamPaginas; i++ {
		posY := recrecQuadrosIn.Y + float32(i)*slotPaginas
		posIni := rl.Vector2{X: recrecQuadrosIn.X, Y: posY}
		posEnd := rl.Vector2{X: recrecQuadrosIn.X + larguraFixa, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}

	for i := 1; i < globais.TamFisica; i++ {
		posY := recrecMemoriaIn.Y + float32(i)*slotPaginas
		posIni := rl.Vector2{X: recrecMemoriaIn.X, Y: posY}
		posEnd := rl.Vector2{X: recrecMemoriaIn.X + larguraFixa, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}
}

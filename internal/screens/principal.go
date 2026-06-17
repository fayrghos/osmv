package screens

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/fayrghos/osmv/internal/state"
	"github.com/fayrghos/osmv/internal/ui"
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type vetorTexto struct {
	valor string
}

var (
	pivo          = -1
	once          sync.Once
	numPaginas    []int
	textosPaginas = []vetorTexto{{"7A4B92C1"}, {"3F8E2D56"}, {"9C104B7E"}, {"D52A8F43"}, {"6B491E2C"}, {"8E3C9A57"}, {"2F5B8D10"}, {"C9A4E6B2"}, {"41D2B8E9"}, {"E07F3C4A"}}
)

func AtualizarPrincipal(globais *state.Globais) {
	once.Do(func() {
		InicializarSimulação(globais)
	})
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
				if pivo < -1 {
					pivo = -1
				} else if pivo > globais.TamPaginas*2 {
					return
				}
				if pivo < globais.TamPaginas {
					pivo++
					globais.Processos[pivo].Pagina = VerificadorSimulacao(globais, pivo)

				}
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

func ConverterBin(valor int) int {
	var valorBin int = 0
	decimal := 1
	var i int = valor
	for i >= 1 {
		valorBin += (i % 2) * decimal
		decimal *= 10
		i /= 2
	}
	return valorBin
}

func InicializarSimulação(globais *state.Globais) {
	globais.Processos = make([]state.BlocoProcesso, 10)
	numPaginas = make([]int, globais.TamPaginas)
	temp := globais.NumBits / 10
	for i := range 10 {
		globais.Processos[i].Texto = textosPaginas[i].valor
		globais.Processos[i].IntervaloBin = ConverterBin(i * temp)
	}
	for j := 0; j < globais.TamPaginas; j++ {
		numPaginas[j] = ConverterBin(j)
	}
}

func VerificadorSimulacao(globais *state.Globais, pivo int) int {
	if pivo < 0 {
		return -1
	}
	bitsParaValidar := globais.TamPaginas / 2
	i := globais.Processos[pivo].IntervaloBin
	for ; i >= int(math.Pow10(bitsParaValidar)); i /= 10 {
	}
	for j := 0; j < globais.TamPaginas; j++ {
		if i == numPaginas[j] {
			return j
		}
	}
	return -1
}

func ExibidorSimulacao(globais *state.Globais, pivo int, rec rl.Rectangle, recTabela rl.Rectangle) {
	if pivo >= 0 && pivo < len(globais.Processos) {
		if globais.Processos[pivo].Pagina >= 0 || globais.Processos[pivo].Pagina < globais.TamPaginas {
			rl.DrawText(globais.Processos[pivo].Texto, int32((rec.X+rec.Width)/2),
				int32(rec.Y*(float32(globais.Processos[pivo].Pagina+1))), 10, rl.White)
		}
		rl.DrawText(strconv.Itoa(globais.Processos[pivo].IntervaloBin), int32((recTabela.X+recTabela.Width)/2),
			int32(recTabela.Y), 10, rl.White)

	}
}

func DesenharPrincipal(globais *state.Globais) {
	// ------------------------------
	// Processos
	// ------------------------------
	recProcessos := rl.Rectangle{
		X:      77,
		Y:      77,
		Width:  230,
		Height: 614,
	}
	rl.DrawRectangleRounded(recProcessos, 0.1, 4, ui.CorPrincipal)

	recProcessosIn := rl.Rectangle{
		X:      88,
		Y:      146,
		Width:  208,
		Height: 533,
	}
	rl.DrawRectangleRec(recProcessosIn, ui.CorSecundaria)

	rl.DrawRectangleRec(
		rl.Rectangle{
			X:      121,
			Y:      146,
			Width:  3,
			Height: 533,
		},
		ui.CorPrincipal,
	)

	utils.DesenharTextoCentro(
		utils.Texto{
			Conteudo: "Processos",
			Tam:      32,
			Fonte:    globais.FonteSans,
		},
		rl.Vector2{X: 192, Y: 111},
		rl.White,
	)

	// ------------------------------
	// Páginas
	// ------------------------------
	recPaginas := rl.Rectangle{
		X:      338,
		Y:      77,
		Width:  230,
		Height: 614,
	}
	rl.DrawRectangleRounded(recPaginas, 0.1, 4, ui.CorPrincipal)

	recPaginasIn := rl.Rectangle{
		X:      348,
		Y:      146,
		Width:  208,
		Height: 533,
	}
	rl.DrawRectangleRec(recPaginasIn, ui.CorSecundaria)

	utils.DesenharTextoCentro(
		utils.Texto{
			Conteudo: "Páginas",
			Tam:      32,
			Fonte:    globais.FonteSans,
		},
		rl.Vector2{X: 452, Y: 111},
		rl.White,
	)

	// ------------------------------
	// Operações
	// ------------------------------
	recOperacoes := rl.Rectangle{
		X:      598,
		Y:      77,
		Width:  268,
		Height: 352,
	}
	rl.DrawRectangleRounded(recOperacoes, 0.08, 4, ui.CorPrincipal)

	recOperacoesIn := rl.Rectangle{
		X:      609,
		Y:      146,
		Width:  245,
		Height: 271,
	}
	rl.DrawRectangleRec(recOperacoesIn, ui.CorSecundaria)

	utils.DesenharTextoCentro(
		utils.Texto{
			Conteudo: "Operações",
			Tam:      32,
			Fonte:    globais.FonteSans,
		},
		rl.Vector2{X: 730, Y: 111},
		rl.White,
	)

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

	recMemoriaIn := rl.Rectangle{
		X:      907,
		Y:      146,
		Width:  371,
		Height: 533,
	}
	rl.DrawRectangleRec(recMemoriaIn, ui.CorSecundaria)

	rl.DrawRectangleRec(
		rl.Rectangle{
			X:      940,
			Y:      146,
			Width:  3,
			Height: 533,
		},
		ui.CorPrincipal,
	)

	utils.DesenharTextoCentro(
		utils.Texto{
			Conteudo: "Memória Física",
			Tam:      32,
			Fonte:    globais.FonteSans,
		},
		rl.Vector2{X: 1093, Y: 111},
		rl.White,
	)

	// ------------------------------
	// Botões
	// ------------------------------
	globais.BotaoContinuar.Desenhar()
	globais.BotaoPasso.Desenhar()
	globais.BotaoVoltar.Desenhar()

	slotPaginas := recPaginasIn.Height / float32(globais.TamPaginas)
	slotPaginasMemo := recMemoriaIn.Height / float32(globais.TamLogica)

	for i := 1; i < globais.TamPaginas; i++ {
		posY := recPaginasIn.Y + float32(i)*slotPaginas
		posIni := rl.Vector2{X: recPaginasIn.X, Y: posY}
		posEnd := rl.Vector2{X: recPaginasIn.X + recProcessosIn.Width, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}

	for i := 1; i < globais.TamLogica; i++ {
		posY := recMemoriaIn.Y + float32(i)*slotPaginasMemo
		posIni := rl.Vector2{X: recMemoriaIn.X, Y: posY}
		posEnd := rl.Vector2{X: recMemoriaIn.X + recMemoriaIn.Width, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}

	centroTabelaPaginas := ((recProcessosIn.X + recProcessosIn.Width) / 2) + 20
	slotTabelaFixa := recProcessosIn.Height / 10
	for i := range 10 {
		posY := recProcessosIn.Y + float32(i)*slotTabelaFixa

		rl.DrawText(textosPaginas[i].valor, int32(centroTabelaPaginas), int32(posY+20), 20, rl.White)

		if i < 9 {
			linhaY := posY + slotTabelaFixa
			posIni := rl.Vector2{X: recProcessosIn.X, Y: linhaY}
			posEnd := rl.Vector2{X: recProcessosIn.X + recProcessosIn.Width, Y: linhaY}
			rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
		}
	}

	ExibidorSimulacao(globais, pivo, recPaginasIn, recOperacoesIn)

	// ------------------------------
	// Texto Bits Centralizar dps
	// ------------------------------

	diviNumBytesMemo := globais.TamFisica / globais.TamLogica
	posicaoTexto := rl.Vector2{X: recMemoria.X, Y: recMemoriaIn.Y + 30}
	yInicial := recMemoriaIn.Y

	for i := 0; i < globais.TamLogica+1; i++ {
		posicaoTexto.Y = yInicial + float32(i)*slotPaginasMemo - 20
		rl.DrawTextEx(*globais.FonteSans, strconv.Itoa(i*diviNumBytesMemo), posicaoTexto, 20, 1, rl.White)
	}

	//Texto tabela
	diviNumBytesPage := globais.NumBits / globais.TamPaginas
	posicaoTextoQuadros := rl.Vector2{X: recPaginasIn.X, Y: recPaginasIn.Y + 30}
	yInicialQuadros := recPaginasIn.Y

	for i := 0; i < globais.TamPaginas+1; i++ {
		posicaoTextoQuadros.Y = yInicialQuadros + float32(i)*slotPaginas
		rl.DrawTextEx(*globais.FonteSans, strconv.Itoa(i*diviNumBytesPage), posicaoTextoQuadros, 20, 1, rl.White)
	}

	//Texto tabela fixa
	diviNumBytesFixo := globais.NumBits / 10
	posicaoTextoFixo := rl.Vector2{X: recProcessosIn.X, Y: recProcessosIn.Y + 30}
	yInicialFixo := recProcessosIn.Y

	for i := range 11 {
		posicaoTextoFixo.Y = yInicialFixo + float32(i)*slotTabelaFixa
		rl.DrawTextEx(*globais.FonteSans, strconv.Itoa(i*diviNumBytesFixo), posicaoTextoFixo, 20, 1, rl.White)
	}
}

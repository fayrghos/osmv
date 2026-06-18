package screens

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/fayrghos/osmv/internal/state"
	"github.com/fayrghos/osmv/internal/tools"
	"github.com/fayrghos/osmv/internal/ui"
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type vetorTexto struct {
	valor string
}

var (
	Memoria    tools.Tabela_Hash
	pivo       = -1
	pivoMemo   = -1
	once       sync.Once
	once2      sync.Once
	once3      sync.Once
	numPaginas []int
	recBaseBin rl.Rectangle = rl.Rectangle{
		X:      609,
		Y:      146,
		Width:  245,
		Height: 271,
	}
	recBusca rl.Rectangle = rl.Rectangle{
		X:      609,
		Y:      146,
		Width:  0,
		Height: 271,
	}
	recMemo rl.Rectangle = rl.Rectangle{
		X:      907,
		Y:      146,
		Width:  371,
		Height: 533,
	}
	textosPaginas = []vetorTexto{{"7A4B92C1"}, {"3F8E2D56"}, {"9C104B7E"}, {"D52A8F43"}, {"6B491E2C"}, {"8E3C9A57"}, {"2F5B8D10"}, {"C9A4E6B2"}, {"41D2B8E9"}, {"E07F3C4A"}}
)

// logica tela principal
func AtualizarPrincipal(globais *state.Globais) {
	recrecQuadrosIn := rl.Rectangle{
		X:      348,
		Y:      146,
		Width:  208,
		Height: 533,
	}
	recrecTabelaIn := rl.Rectangle{
		X:      609,
		Y:      146,
		Width:  245,
		Height: 271,
	}

	once.Do(func() {
		InicializarSimulação(globais)
	})
	globais.InicializarTela(state.TelaPrincipal, func() {
		globais.BotaoContinuar = ui.Botao{
			Pos: rl.Vector2{X: 598, Y: 441},
			Tam: rl.Vector2{X: 268, Y: 75},
			Rotulo: utils.Texto{
				Conteudo: "Avançar",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: false,
			Clique: func() {
				once2.Do(func() {
					soma(globais)
				})
				globais.BotaoVoltar.Travado = false
				globais.BotaoReiniciar.Travado = false

				if pivo < -1 {
					pivo = -1
				}
				if pivoMemo < -1 {
					pivoMemo = -1
				}

				if pivo < len(globais.Processos)-1 {
					pivo++
					globais.Processos[pivo].Pagina = VerificadorSimulacao(globais, pivo)
					globais.Processos[pivo].CoordBin = recrecTabelaIn
					globais.Processos[pivo].CoordText = recrecQuadrosIn
				} else {
					once3.Do(func() {
						Memoria.Preencher(globais)
					})
					if pivoMemo < globais.TamLogica-1 {
						pivoMemo++
						if pivoMemo == globais.TamLogica-1 {
							globais.BotaoContinuar.Travado = true
						}
						Memoria.TabelaMemo[pivoMemo].Coordenada = recMemo
					}
				}
			},
		}

		globais.BotaoVoltar = ui.Botao{
			Pos: rl.Vector2{X: 598, Y: 528},
			Tam: rl.Vector2{X: 268, Y: 75},
			Rotulo: utils.Texto{
				Conteudo: "Voltar",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: true,
			Clique: func() {
				globais.BotaoContinuar.Travado = false

				if pivo < -1 {
					pivo = -1
				}
				if pivo >= 0 && pivoMemo <= -1 {
					pivo--
				} else if pivoMemo >= 0 {
					pivoMemo--
				}

				if pivoMemo < 0 && pivo < 0 {
					globais.BotaoVoltar.Travado = true
					globais.BotaoReiniciar.Travado = true
				}
			},
		}

		globais.BotaoReiniciar = ui.Botao{
			Pos: rl.Vector2{X: 598, Y: 616},
			Tam: rl.Vector2{X: 268, Y: 75},
			Rotulo: utils.Texto{
				Conteudo: "Reiniciar",
				Tam:      32,
				Fonte:    globais.FonteSans,
			},
			Travado: true,
			Clique: func() {
				for !globais.BotaoVoltar.Travado {
					globais.BotaoVoltar.Clique()
				}
			},
		}
	})

	globais.BotaoContinuar.Atualizar()
	globais.BotaoReiniciar.Atualizar()
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
	if pivo < 0 || globais.TamPaginas <= 0 {
		return -1
	}

	bitsParaValidar := int(math.Log2(float64(globais.TamPaginas)))
	if bitsParaValidar < 1 {
		bitsParaValidar = 1
	}

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

func 固定(globais *state.Globais) {
	bitsParaValidar := float32(math.Log2(float64(globais.TamPaginas)))
	if bitsParaValidar < 1 {
		bitsParaValidar = 1
	}
	recBusca.X = 615
	recBusca.Y = 156
	recBusca.Width = bitsParaValidar*12 + 4
	recBusca.Height = recBaseBin.Height - 20
}

func soma(globais *state.Globais) {
	bitsParaValidar := float32(math.Log2(float64(globais.TamPaginas)))
	if bitsParaValidar < 1 {
		bitsParaValidar = 1
	}
	recBusca.X = 615
	recBusca.Y = 156
	recBusca.Width = bitsParaValidar*12 + 4
	recBusca.Height = recBaseBin.Height - 20
}

func ExibidorMemoria(globais *state.Globais, i int) {
	if i < 0 || i >= globais.TamLogica {
		return
	}

	proc := Memoria.TabelaMemo[i]

	if proc.Valor >= 0 && proc.Valor < globais.TamPaginas {
		slotAlturaQuadros := proc.Coordenada.Height / float32(globais.TamLogica)
		yQuadro := proc.Coordenada.Y + (float32(proc.Valor) * slotAlturaQuadros) + (slotAlturaQuadros / 2) - 10
		xQuadro := proc.Coordenada.X + 45

		colisao := 0
		for j := 0; j < i; j++ {
			if Memoria.TabelaMemo[j].Valor == proc.Valor {
				colisao++
			}
		}

		if colisao > 0 {
			deslocamentoX := float32(colisao%2) * 95
			deslocamentoY := float32(colisao/2) * 22
			xQuadro += deslocamentoX
			yQuadro += deslocamentoY
			if colisao >= 2 {
				yQuadro -= 6
			}
		}

		rl.DrawTextEx(
			*globais.FonteSans,
			proc.Nome,
			rl.Vector2{X: xQuadro, Y: yQuadro},
			20,
			1,
			rl.White,
		)
	}
}

func ExibidorSimulacao(globais *state.Globais, i int) {
	if i < 0 || i >= len(globais.Processos) {
		return
	}

	proc := globais.Processos[i]
	alturaLinhaOperacao := float32(25)
	yOperacao := proc.CoordBin.Y + 15 + (float32(i) * alturaLinhaOperacao)
	xOperacao := proc.CoordBin.X + 15
	rl.DrawTextEx(
		*globais.FonteSans,
		strconv.Itoa(proc.IntervaloBin),
		rl.Vector2{X: xOperacao, Y: yOperacao},
		20,
		10,
		rl.White,
	)

	if proc.Pagina >= 0 && proc.Pagina < globais.TamPaginas {
		slotAlturaQuadros := proc.CoordText.Height / float32(globais.TamPaginas)

		// Centralização base vertical ajustada para o tamanho 16
		yQuadro := proc.CoordText.Y + (float32(proc.Pagina) * slotAlturaQuadros) + (slotAlturaQuadros / 2) - 8

		// Margem esquerda estável para alinhar a grade perfeitamente
		xQuadro := proc.CoordText.X + 50

		colisao := 0
		for j := 0; j < i; j++ {
			if globais.Processos[j].Pagina == proc.Pagina {
				colisao++
			}
		}

		// CORREÇÃO: Multiplicadores expandidos (80px horizontal / 18px vertical) para dar o respiro necessário
		if colisao > 0 {
			deslocamentoX := float32(colisao%2) * 80
			deslocamentoY := float32(colisao/2) * 18
			xQuadro += deslocamentoX
			yQuadro += deslocamentoY

			if colisao >= 2 {
				yQuadro -= 5
			}
		}

		rl.DrawTextEx(*globais.FonteSans, proc.Texto, rl.Vector2{X: xQuadro, Y: yQuadro}, 16, 1, rl.White)
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

	rl.DrawRectangleRec(
		rl.Rectangle{
			X:      381,
			Y:      146,
			Width:  3,
			Height: 533,
		},
		ui.CorPrincipal,
	)

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
	globais.BotaoReiniciar.Desenhar()
	globais.BotaoVoltar.Desenhar()

	slotPaginas := recPaginasIn.Height / float32(globais.TamPaginas)
	slotPaginasMemo := recMemoriaIn.Height / float32(globais.TamLogica)

	for i := 1; i < globais.TamPaginas; i++ {
		posY := recPaginasIn.Y + float32(i)*slotPaginas
		posIni := rl.Vector2{X: recPaginasIn.X, Y: posY}
		posEnd := rl.Vector2{X: recPaginasIn.X + recPaginasIn.Width, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}

	for i := 1; i < globais.TamLogica; i++ {
		posY := recMemoriaIn.Y + float32(i)*slotPaginasMemo
		posIni := rl.Vector2{X: recMemoriaIn.X, Y: posY}
		posEnd := rl.Vector2{X: recMemoriaIn.X + recMemoriaIn.Width, Y: posY}
		rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
	}

	centroTabelaPaginas := (recProcessosIn.X + recProcessosIn.Width) / 2
	slotTabelaFixa := recProcessosIn.Height / 10
	for i := range 10 {
		posY := recProcessosIn.Y + float32(i)*slotTabelaFixa

		rl.DrawTextEx(
			*globais.FonteSans,
			textosPaginas[i].valor,
			rl.Vector2{X: centroTabelaPaginas + 5, Y: posY + 15},
			24,
			1,
			rl.White,
		)

		if i < 9 {
			linhaY := posY + slotTabelaFixa
			posIni := rl.Vector2{X: recProcessosIn.X, Y: linhaY}
			posEnd := rl.Vector2{X: recProcessosIn.X + recProcessosIn.Width, Y: linhaY}
			rl.DrawLineEx(posIni, posEnd, 3, ui.CorPrincipal)
		}
	}

	//Chamada da função para exibir a simulação
	for i := 0; i <= pivo; i++ {
		ExibidorSimulacao(globais, i)
	}
	for j := 0; j <= pivoMemo; j++ {
		ExibidorMemoria(globais, j)
	}

	// ------------------------------
	// Texto Bits
	// ------------------------------

	diviNumBytesMemo := globais.TamFisica / globais.TamLogica
	yInicial := recMemoriaIn.Y

	for i := range globais.TamLogica {
		posY := yInicial + float32(i)*slotPaginasMemo + (slotPaginasMemo / 2) - 10
		rl.DrawTextEx(*globais.FonteSans, fmt.Sprintf("%02d", i*diviNumBytesMemo), rl.Vector2{X: recMemoria.X + 12, Y: posY}, 20, 1, rl.White)
	}

	//Texto tabela
	diviNumBytesPage := globais.NumBits / globais.TamPaginas
	yInicialQuadros := recPaginasIn.Y

	for i := range globais.TamPaginas {
		posY := yInicialQuadros + float32(i)*slotPaginas + (slotPaginas / 2) - 10
		rl.DrawTextEx(*globais.FonteSans, fmt.Sprintf("%02d", i*diviNumBytesPage), rl.Vector2{X: recPaginasIn.X + 6, Y: posY}, 20, 1, rl.White)
	}

	//Texto tabela fixa
	diviNumBytesFixo := globais.NumBits / 10
	posicaoTextoFixo := rl.Vector2{X: recProcessosIn.X + 8, Y: recProcessosIn.Y + 30}
	yInicialFixo := recProcessosIn.Y

	for i := range 10 {
		posicaoTextoFixo.Y = yInicialFixo + float32(i)*slotTabelaFixa + 17
		rl.DrawTextEx(*globais.FonteSans, fmt.Sprintf("%.2d", i*diviNumBytesFixo), posicaoTextoFixo, 20, 1, rl.White)
	}
}

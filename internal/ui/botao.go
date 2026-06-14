package ui

import (
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Botões delicisiosamente clicáveis
type Botao struct {
	Pos     rl.Vector2
	Tam     rl.Vector2
	Rotulo  utils.Texto
	Travado bool
	Clique  func()

	destacado     bool
	stampDestaque float64
}

// Redesenha o botão
func (bot Botao) Desenhar() {
	recBox := rl.Rectangle{
		X:      bot.Pos.X,
		Y:      bot.Pos.Y,
		Width:  bot.Tam.X,
		Height: bot.Tam.Y,
	}

	corDesenhada := rl.White

	switch {
	case bot.Travado:
		corDesenhada.A = 80
	case bot.stampDestaque > rl.GetTime():
		rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 4, CorSelecao)
	case bot.destacado:
		rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 3, CorDestaque)
	}
	rl.DrawRectangleRounded(recBox, 0.5, 1, CorPrincipal)

	medidaRotulo := rl.MeasureTextEx(*bot.Rotulo.Fonte, bot.Rotulo.Conteudo, bot.Rotulo.Tam, 1)
	rl.DrawTextEx(
		*bot.Rotulo.Fonte,
		bot.Rotulo.Conteudo,
		bot.Pos.Add(bot.Tam.Subtract(medidaRotulo).Divide(rl.Vector2{X: 2, Y: 2})),
		bot.Rotulo.Tam,
		1,
		corDesenhada,
	)
}

// Atualiza a lógica do botão
func (bot *Botao) Atualizar() {
	retanColisao := rl.Rectangle{
		X:      bot.Pos.X,
		Y:      bot.Pos.Y,
		Width:  bot.Tam.X,
		Height: bot.Tam.Y,
	}

	if !bot.Travado && rl.CheckCollisionPointRec(rl.GetMousePosition(), retanColisao) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			bot.stampDestaque = rl.GetTime() + 0.1
			bot.Clique()
		}
		bot.destacado = true
	} else {
		bot.destacado = false
	}
}

package ui

import (
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Caixas de erro
type Errobox struct {
	Pos   rl.Vector2
	Tam   rl.Vector2
	Campo utils.Texto
}

func (box Errobox) Desenhar() {
	if box.Campo.Conteudo == "" {
		return
	}

	recBox := rl.Rectangle{
		X:      box.Pos.X,
		Y:      box.Pos.Y,
		Width:  box.Tam.X,
		Height: box.Tam.Y,
	}

	rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 4, CorErro)
	rl.DrawRectangleRounded(recBox, 0.5, 1, CorPrincipal)

	medidaCampo := rl.MeasureTextEx(*box.Campo.Fonte, box.Campo.Conteudo, box.Campo.Tam, 1)
	rl.DrawTextEx(
		*box.Campo.Fonte,
		box.Campo.Conteudo,
		box.Pos.Add(box.Tam.Subtract(medidaCampo).Divide(rl.Vector2{X: 2, Y: 2})),
		box.Campo.Tam,
		1,
		CorErro,
	)
}

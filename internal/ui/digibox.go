package ui

import (
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Digibox struct {
	Pos         rl.Vector2
	Tam         rl.Vector2
	Texto       string
	Selecionada bool
	Destacada   bool
	Fonte       *rl.Font
	FonteTam    float32
}

func (box Digibox) Desenhar() {
	recBox := rl.Rectangle{
		X:      box.Pos.X,
		Y:      box.Pos.Y,
		Width:  box.Tam.X,
		Height: box.Tam.Y,
	}

	if box.Selecionada {
		rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 3, CorSelecao)
	} else if box.Destacada {
		rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 2, CorDestaque)
	}
	rl.DrawRectangleRounded(recBox, 0.5, 1, CorPrincipal)

	textoDesenhar := ""
	if len(box.Texto) > 0 {
		textoDesenhar = box.Texto
	} else if box.Selecionada && int(rl.GetTime()*4)%2 == 0 {
		textoDesenhar = "_"
	}

	medida := rl.MeasureTextEx(*box.Fonte, textoDesenhar, box.FonteTam, 1)
	rl.DrawTextEx(
		*box.Fonte,
		textoDesenhar,
		box.Pos.Add(box.Tam.Subtract(medida).Divide(rl.Vector2{X: 2, Y: 2})),
		box.FonteTam,
		1,
		rl.White,
	)
}

func (box *Digibox) ReceberNums() {
	digitado := rl.GetCharPressed()

	if digitado >= '\u0030' && digitado <= '\u0039' && len(box.Texto) < 5 {
		box.Texto += string(digitado)
	} else if utils.IsKeyPressedDouble(rl.KeyBackspace) && len(box.Texto) > 0 {
		box.Texto = box.Texto[:len(box.Texto)-1]
	}
}

func (box *Digibox) Atualizar() {
	retanColisao := rl.Rectangle{
		X:      box.Pos.X,
		Y:      box.Pos.Y,
		Width:  box.Tam.X,
		Height: box.Tam.Y,
	}

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), retanColisao) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			box.Selecionada = true
		}
		box.Destacada = true
	} else {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			box.Selecionada = false
		}
		box.Destacada = false
	}

	if box.Selecionada {
		box.ReceberNums()
	}
}

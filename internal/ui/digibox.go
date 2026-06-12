package ui

import (
	"github.com/fayrghos/osmv/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Caixas de digitação
type Digibox struct {
	Pos    rl.Vector2
	Tam    rl.Vector2
	Campo  utils.Texto
	Titulo utils.Texto

	selecionada bool
	destacada   bool
}

// Redesenha a caixa
func (box Digibox) Desenhar() {
	recBox := rl.Rectangle{
		X:      box.Pos.X,
		Y:      box.Pos.Y,
		Width:  box.Tam.X,
		Height: box.Tam.Y,
	}

	if box.selecionada {
		rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 4, CorSelecao)
	} else if box.destacada {
		rl.DrawRectangleRoundedLinesEx(recBox, 0.5, 1, 3, CorDestaque)
	}
	rl.DrawRectangleRounded(recBox, 0.5, 1, CorPrincipal)

	textoDesenhar := ""
	if len(box.Campo.Conteudo) > 0 {
		textoDesenhar = box.Campo.Conteudo
	} else if box.selecionada && int(rl.GetTime()*4)%2 == 0 {
		textoDesenhar = "_"
	}

	medidaCampo := rl.MeasureTextEx(*box.Campo.Fonte, textoDesenhar, box.Campo.Tam, 1)
	rl.DrawTextEx(
		*box.Campo.Fonte,
		textoDesenhar,
		box.Pos.Add(box.Tam.Subtract(medidaCampo).Divide(rl.Vector2{X: 2, Y: 2})),
		box.Campo.Tam,
		1,
		rl.White,
	)

	medidaTitulo := rl.MeasureTextEx(*box.Titulo.Fonte, box.Titulo.Conteudo, box.Titulo.Tam, 1)
	rl.DrawTextEx(
		*box.Titulo.Fonte,
		box.Titulo.Conteudo,
		box.Pos.Add(rl.Vector2{X: 3, Y: -medidaTitulo.Y - 10}),
		box.Titulo.Tam,
		1,
		rl.White,
	)
}

// Permite que o usuário insira números na caixa
func (box *Digibox) ReceberNums() {
	digitado := rl.GetCharPressed()

	if digitado >= '0' && digitado <= '9' && len(box.Campo.Conteudo) < 5 {
		box.Campo.Conteudo += string(digitado)
	} else if utils.IsKeyPressedDouble(rl.KeyBackspace) && len(box.Campo.Conteudo) > 0 {
		box.Campo.Conteudo = box.Campo.Conteudo[:len(box.Campo.Conteudo)-1]
	}
}

// Atualiza a lógica da caixa
func (box *Digibox) Atualizar() {
	retanColisao := rl.Rectangle{
		X:      box.Pos.X,
		Y:      box.Pos.Y,
		Width:  box.Tam.X,
		Height: box.Tam.Y,
	}

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), retanColisao) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			box.selecionada = true
		}
		box.destacada = true
	} else {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			box.selecionada = false
		}
		box.destacada = false
	}

	if box.selecionada {
		box.ReceberNums()
	}
}

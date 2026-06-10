package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO: Botar no globs
var botaoSelecionado *Botao

type Botao struct {
	Pos       rl.Vector2
	Tamanho   rl.Vector2
	Destacado bool
	Texto     string
	CliqueEsq func(b *Botao)
	Fonte     *rl.Font
}

func (b *Botao) ReagirCursor() {
	retanColisao := rl.Rectangle{
		X:      b.Pos.X,
		Y:      b.Pos.Y,
		Width:  b.Tamanho.X,
		Height: b.Tamanho.Y,
	}

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), retanColisao) {
		b.Destacado = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			botaoSelecionado = b
		}
	} else {
		b.Destacado = false
	}
}

func (b *Botao) ReceberDigitos() {
	digitado := rl.GetCharPressed()
	if digitado >= 48 && digitado <= 57 && len(b.Texto) < 5 {
		b.Texto += string(digitado)
	} else if rl.IsKeyPressed(rl.KeyBackspace) {
		b.Texto = b.Texto[:len(b.Texto)-1]
	}
}

type GrupoBotoes []Botao

func (gb GrupoBotoes) DesenharTodos() {
	for i := range gb {
		if &gb[i] == botaoSelecionado {
			rl.DrawRectangleV(gb[i].Pos.AddValue(-3), gb[i].Tamanho.AddValue(6), CorSelecao)
		} else if gb[i].Destacado {
			rl.DrawRectangleV(gb[i].Pos.AddValue(-3), gb[i].Tamanho.AddValue(6), CorDestaque)
		}
		rl.DrawRectangleV(gb[i].Pos, gb[i].Tamanho, CorPrincipal)

		rl.DrawTextEx(*gb[i].Fonte, gb[i].Texto, gb[i].Pos, 32, 1, rl.White)
	}
}

func (gb GrupoBotoes) AtualizarTodos() {
	for i := range gb {
		gb[i].ReagirCursor()

		if &gb[i] == botaoSelecionado {
			gb[i].ReceberDigitos()
		}
	}
}

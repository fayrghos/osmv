package utils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Armazena facilmente objetos comuns de texto
type Texto struct {
	Conteudo string
	Tam      float32
	Fonte    *rl.Font
}

// Carrega a fonte corretamente e melhora a resolução dela
func CarregarFonte(caminho string, tam int32) rl.Font {
	alfas := " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	acentos := "áéíóúâêîôûãõÁÉÍÓÚÂÊÎÔÛÃÕçÇ"
	pontuacao := ",.:?!_"

	caracteres := []rune(alfas + acentos + pontuacao)
	saida := rl.LoadFontEx(caminho, tam*3, caracteres, int32(len(caracteres)))

	rl.SetTextureFilter(saida.Texture, rl.FilterBilinear)
	return saida
}

// É tipo o IsKeyDown(), mas apropriado pra digitação
func IsKeyPressedDouble(key int32) bool {
	return rl.IsKeyPressed(key) || rl.IsKeyPressedRepeat(key)
}

// Desenha um texto centralizado em 'pos'
func DesenharTextoCentro(texto Texto, pos rl.Vector2, cor color.RGBA) {
	medidaTexPaginas := rl.MeasureTextEx(*texto.Fonte, texto.Conteudo, texto.Tam, 1)
	rl.DrawTextEx(
		*texto.Fonte,
		texto.Conteudo,
		pos.Subtract(medidaTexPaginas.Scale(0.5)),
		texto.Tam,
		1,
		cor,
	)
}

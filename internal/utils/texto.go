package utils

import rl "github.com/gen2brain/raylib-go/raylib"

// Armazena facilmente objetos comuns de texto
type Texto struct {
	Conteudo string
	Tam      float32
	Fonte    *rl.Font
}

// Carrega a fonte e melhora a resolução dela
func CarregarFonte(caminho string, tam int32) rl.Font {
	saida := rl.LoadFontEx(caminho, tam*3, nil)
	rl.SetTextureFilter(saida.Texture, rl.FilterBilinear)
	return saida
}

// É tipo o IsKeyDown(), mas apropriado pra digitação
func IsKeyPressedDouble(key int32) bool {
	return rl.IsKeyPressed(key) || rl.IsKeyPressedRepeat(key)
}

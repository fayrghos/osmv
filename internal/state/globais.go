package state

import (
	"github.com/fayrghos/osmv/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Tamanho da janela
const (
	Larg = 1366
	Altu = 768
)

// Enum das telas
type Tela int

const (
	TelaInicial Tela = iota
	TelaPrincipal
)

// Como é bom te ver de novo, Globs!
type Globais struct {
	FonteSans *rl.Font
	BoxErro   ui.Errobox

	TelaAtual      Tela
	telasIniciadas map[Tela]bool

	TamPaginas    int
	BoxTamPaginas ui.Digibox
	TamFisica     int
	BoxTamFisica  ui.Digibox
	TamLogica     int
	BoxTamLogica  ui.Digibox
}

// Essa função roda algo exatamente uma vez quando uma tela inicia
func (gs *Globais) InicializarTela(tela Tela, funcaoInicio func()) {
	if gs.telasIniciadas == nil {
		gs.telasIniciadas = make(map[Tela]bool, 10) // Máximo de telas
	}

	if gs.telasIniciadas[tela] == false {
		gs.telasIniciadas[tela] = true
		funcaoInicio()
	}
}

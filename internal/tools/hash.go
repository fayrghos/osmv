package tools

import (
	"strconv"

	"github.com/fayrghos/osmv/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pagina struct {
	Valor      int
	Nome       string
	Processos  []state.BlocoProcesso
	Coordenada rl.Rectangle
}

type Tabela_Hash struct {
	TabelaMemo     []Pagina
	TabelaPreOrdem []Pagina
}

func (t *Tabela_Hash) Preencher(globais *state.Globais) {
	t.TabelaMemo = make([]Pagina, globais.TamLogica)
	t.TabelaPreOrdem = make([]Pagina, globais.TamPaginas)

	for i := 0; i < globais.TamPaginas; i++ {
		t.TabelaPreOrdem[i].Valor = i
		t.TabelaPreOrdem[i].Processos = make([]state.BlocoProcesso, 0)

		for j := 0; j < len(globais.Processos); j++ {
			if globais.Processos[j].Pagina == t.TabelaPreOrdem[i].Valor {
				t.TabelaPreOrdem[i].Nome = ("Page: " + strconv.Itoa(i))
				t.TabelaPreOrdem[i].Processos = append(t.TabelaPreOrdem[i].Processos, globais.Processos[j])
			}
		}
	}
	t.Hash(globais)
}

func (tab *Tabela_Hash) Hash(globais *state.Globais) {
	tab.TabelaMemo[0].Nome = "Espaço SO"
	tab.TabelaMemo[0].Valor = 0

	for i := 0; i < globais.TamPaginas; i++ {
		x := tab.TabelaPreOrdem[i].Valor % globais.TamLogica
		if x != 0 {
			tab.TabelaMemo[x] = tab.TabelaPreOrdem[i]
		}
	}
}

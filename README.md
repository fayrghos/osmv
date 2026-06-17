# OSMV

Um pequeno visualizador de memória escrito em Go, em conjunto da biblioteca Raylib para a disciplina de Sistemas Operacionais.

## Navegação

O programa foi projeto para operar utilizando o mouse como principal ferramenta. O teclado é reservado apenas a digitação nas caixas de texto.

## Estrutura

- `bin/`: Contém os executáveis após compilações.
- `assets/`: Contém mídia em geral a ser exibida no programa.
- `internal/`: Contém toda a lógica interna do programa.

## Desenvolvimento

Para desenvolver o projeto, é necessário ter um compilador de Go, o Make e a biblioteca raylib-go instalados em sua máquina. Qualquer ambiente de desenvolvimento pode ser utilizado.

- [Instalação do Go](https://go.dev/doc/install)
- [Instalação da raylib-go](https://github.com/gen2brain/raylib-go#requirements)

## Compilando

Após ter os pré-requisitos instalados, basta apenas rodar as regras do make na raíz do projeto.

### Normal

Para uma compilação casual do programa, basta chamar o make sem nenhum argumento. O programa será executado automaticamente em seguida.

```bash
make
```

### Rápido

Caso você queira que os valores do programa sejam preenchidos automaticamente, basta executar a regra de execução rápida.

```bash
make rapido # "make r" também é aceito.
```

### Debug

Caso o programa esteja com algum bug estranho e você queira fazer um debug, basta executar a regra de debug.

```bash
make debug # "make d" também é aceito.
```

## Apêndice

![QueOSMV](./assets/repo/queosmv.jpeg)

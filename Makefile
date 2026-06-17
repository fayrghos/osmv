BIN := ./bin/osmv
ARGS :=
FLAGS :=


# Roda casualmente o programa
normal: compilar
	./$(BIN) $(ARGS)


# Executa o programa no modo rápido
rapido r: ARGS += --rapido
rapido r: compilar
	./$(BIN) $(ARGS)


# Chama o GDB para fazer debug
debug d: FLAGS += all=-N -l
debug d: compilar
	gdb $(BIN) -q


# Apenas compila o binário
compilar:
	@mkdir -p bin
	go build -gcflags="$(FLAGS)" -tags x11 -o $(BIN) .


.PHONY: all rapido r debug d compilar

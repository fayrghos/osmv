BIN := ./bin/osmv
FLAGS :=


# Roda casualmente o programa
normal: compilar
	./$(BIN)


# Chama o GDB para fazer debug
debug d: FLAGS += all=-N -l
debug d: compilar
	gdb $(BIN) -q


# Apenas compila o binário
compilar:
	@mkdir -p bin
	go build -gcflags="$(FLAGS)" -tags x11 -o $(BIN) .


.PHONY: all debug d compilar

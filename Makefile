BIN := ./bin/osmv


# Roda casualmente o programa
normal: compilar
	./$(BIN)


# Chama o GDB para fazer debug
debug d: compilar
	gdb $(BIN) -q


# Apenas compila o binário
compilar:
	@mkdir -p bin
	go build -gcflags="all=-N -l" -o $(BIN) .


.PHONY: all debug d compilar

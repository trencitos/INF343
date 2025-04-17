# Makefile para Gobierno Mundial
NAME = gobierno
PORT = 50052

docker-gobierno: build run

build:
	docker build -t $(NAME) .

run:
	docker run -p $(PORT):$(PORT) --name $(NAME)_container $(NAME)

stop:
	docker stop $(NAME)_container || true
	docker rm $(NAME)_container || true

clean: stop
	docker rmi $(NAME) || true

help:
	@echo "Comandos disponibles:"
	@echo "  build  - Construir la imagen de Docker"
	@echo "  run    - Ejecutar el contenedor"
	@echo "  stop   - Detener y eliminar el contenedor"
	@echo "  clean  - Limpiar la imagen de Docker"
	@echo "  help   - Mostrar esta ayuda"

.PHONY: build run stop clean help

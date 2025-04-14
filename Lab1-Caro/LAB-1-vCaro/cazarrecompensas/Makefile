# Variables
APP_NAME = cazarrecompensas
PROTO_DIR = ../proto
PROTO_FILE = $(PROTO_DIR)/sistema.proto
OUT_DIR = ../sistema_pb

# Comandos por defecto
all: build

# Generar archivos gRPC desde el .proto
proto:
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_FILE)

build:
	go build -o $(APP_NAME) main.go

run:
	go run main.go

clean:
	rm -f $(APP_NAME)

rebuild: clean proto build

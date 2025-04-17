package main

import (
	"context"
	"log"
	"fmt"
	"math/rand"
	"net"
	"time"

	pb "submundo/proto/grpc-server/proto" // Importa el paquete generado por gRPC
	"google.golang.org/grpc"
)

// Implementación del servidor Submundo
type servidorSubmundo struct {
	pb.UnimplementedSubmundoServiceServer
}

// Lógica para comprar pirata
func (s *servidorSubmundo) ComprarPirata(ctx context.Context, pirata *pb.PirataCapturado) (*pb.ResultadoCompra, error) {
	log.Printf("Intentando vender pirata: %s con recompensa de %d berries", pirata.Nombre, pirata.Recompensa)
	rand.Seed(time.Now().UnixNano())
	resultado := &pb.ResultadoCompra{
		Nombre:     pirata.Nombre,
		Estado:     "Capturado",
		Entrega: "submundo",
	}
	// Verificamos probabilidad de fraude (35%)
	if rand.Float32() < 0.35 {
		log.Println("¡Fraude! El cazarrecompensas no recibió pago.")
		resultado.Caso = "Fraude"
		resultado.MontoPagado = 0
		resultado.Mensaje = "¡Fraude! El cazarrecompensas no recibió pago."

	} else {
		multiplicador := 1.0 + rand.Float32()*0.5 // entre 1.0 y 1.5
		monto := int32(float32(pirata.Recompensa) * multiplicador)

		log.Println("Compra exitosa. El pirata fue adquirido.")
		resultado.Caso = "Exito"
		resultado.MontoPagado = monto
		resultado.Mensaje = fmt.Sprintf("Compra exitosa. El pirata fue vendido por %d berries.", monto)
	}

	return resultado, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Iniciando submundo...")
	
	// Crear un listener para el servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Servidor falla al escuchar: %v", err)
	}
	// Crear un servidor gRPC
	s := grpc.NewServer()
	pb.RegisterSubmundoServiceServer(s, &servidorSubmundo{})
	log.Println("Servidor gRPC del Submundo escuchando en el puerto 50051...")
	// Iniciar el servidor gRPC
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
	}
}

package main

import (
	"context"
	"log"
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
	log.Printf("Intentando comprar pirata: %s con recompensa %d berries", pirata.Nombre, pirata.Recompensa)

	resultado := &pb.ResultadoCompra{}

	// Verificamos probabilidad de fraude (35%)
	if rand.Float32() < 0.35 {
		resultado.Exito = false
		resultado.Mensaje = "¡Fraude! El cazarrecompensas no recibió pago."
		resultado.MontoPagado = 0
		resultado.FueFraude = true
	} else {
		multiplicador := 1.0 + rand.Float32()*0.5 // entre 1.0 y 1.5
		monto := int32(float32(pirata.Recompensa) * multiplicador)

		resultado.Exito = true
		resultado.Mensaje = "Compra exitosa. El pirata fue adquirido."
		resultado.MontoPagado = monto
		resultado.FueFraude = false
	}

	return resultado, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Iniciando submundo...")
	
	// Crear un listener para el servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}

	// Crear un servidor gRPC
	s := grpc.NewServer()
	// Aquí puedes registrar tus servicios en el servidor gRPC
	// Ejemplo: pb.RegisterYourServiceServer(grpcServer, &YourService{})
	pb.RegisterSubmundoServiceServer(s, &servidorSubmundo{})

	log.Println("Servidor gRPC del Submundo escuchando en el puerto 50051...")
	// Iniciar el servidor gRPC
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
	}
}
